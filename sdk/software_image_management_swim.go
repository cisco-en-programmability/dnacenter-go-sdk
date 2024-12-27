package dnac

import (
	"fmt"
	"io"
	"net/http"
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
type ReturnsListOfSoftwareImagesQueryParams struct {
	SiteID                       string  `url:"siteId,omitempty"`                       //Site identifier to get the list of all available products under the site. The default value is the global site.  See https://developer.cisco.com/docs/dna-center/get-site for `siteId`
	ProductNameOrdinal           float64 `url:"productNameOrdinal,omitempty"`           //The product name ordinal is a unique value for each network device product. The productNameOrdinal can be obtained from the response of API `/dna/intent/api/v1/siteWiseProductNames`
	SupervisorProductNameOrdinal float64 `url:"supervisorProductNameOrdinal,omitempty"` //The supervisor engine module ordinal is a unique value for each supervisor module. The `supervisorProductNameOrdinal` can be obtained from the response of API `/dna/intent/api/v1/siteWiseProductNames`
	Imported                     bool    `url:"imported,omitempty"`                     //When the value is set to `true`, it will include physically imported images. Conversely, when the value is set to `false`, it will include image records from the cloud. The identifier for cloud images can be utilized to download images from Cisco.com to the disk.
	Name                         string  `url:"name,omitempty"`                         //Filter with software image or add-on name. Supports partial case-insensitive search. A minimum of 3 characters is required for the search.
	Version                      string  `url:"version,omitempty"`                      //Filter with image version. Supports partial case-insensitive search. A minimum of 3 characters is required for the search.
	Golden                       bool    `url:"golden,omitempty"`                       //When set to `true`, it will retrieve the images marked as tagged golden. When set to `false`, it will retrieve the images marked as not tagged golden.
	Integrity                    string  `url:"integrity,omitempty"`                    //Filter with verified images using Integrity Verification Available values: UNKNOWN, VERIFIED
	HasAddonImages               bool    `url:"hasAddonImages,omitempty"`               //When set to `true`, it will retrieve the images which have add-on images. When set to `false`, it will retrieve the images which do not have add-on images.
	IsAddonImages                bool    `url:"isAddonImages,omitempty"`                //When set to `true`, it will retrieve the images that an add-on image.  When set to `false`, it will retrieve the images that are not add-on images
	Offset                       float64 `url:"offset,omitempty"`                       //The first record to show for this page; the first record is numbered 1. The minimum value is 1.
	Limit                        float64 `url:"limit,omitempty"`                        //The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively.
}
type ReturnsCountOfSoftwareImagesQueryParams struct {
	SiteID                       string  `url:"siteId,omitempty"`                       //Site identifier to get the list of all available products under the site. The default value is the global site.  See https://developer.cisco.com/docs/dna-center/get-site for siteId
	ProductNameOrdinal           float64 `url:"productNameOrdinal,omitempty"`           //The product name ordinal is a unique value for each network device product. The productNameOrdinal can be obtained from the response of the API `/dna/intent/api/v1/siteWiseProductNames`.
	SupervisorProductNameOrdinal float64 `url:"supervisorProductNameOrdinal,omitempty"` //The supervisor engine module ordinal is a unique value for each supervisor module. The `supervisorProductNameOrdinal` can be obtained from the response of API `/dna/intent/api/v1/siteWiseProductNames`
	Imported                     bool    `url:"imported,omitempty"`                     //When the value is set to `true`, it will include physically imported images. Conversely, when the value is set to `false`, it will include image records from the cloud. The identifier for cloud images can be utilised to download images from Cisco.com to the disk.
	Name                         string  `url:"name,omitempty"`                         //Filter with software image or add-on name. Supports partial case-insensitive search. A minimum of 3 characters is required for the search
	Version                      string  `url:"version,omitempty"`                      //Filter with image version. Supports partial case-insensitive search. A minimum of 3 characters is required for the search
	Golden                       string  `url:"golden,omitempty"`                       //When set to `true`, it will retrieve the images marked tagged golden. When set to `false`, it will retrieve the images marked not tagged golden.
	Integrity                    string  `url:"integrity,omitempty"`                    //Filter with verified images using Integrity Verification Available values: UNKNOWN, VERIFIED
	HasAddonImages               bool    `url:"hasAddonImages,omitempty"`               //When set to `true`, it will retrieve the images which have add-on images. When set to `false`, it will retrieve the images which do not have add-on images.
	IsAddonImages                bool    `url:"isAddonImages,omitempty"`                //When set to `true`, it will retrieve the images that an add-on image.  When set to `false`, it will retrieve the images that are not add-on images
}
type RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageQueryParams struct {
	ProductName string  `url:"productName,omitempty"` //Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters is required for the search.
	ProductID   string  `url:"productId,omitempty"`   //Filter with product ID (PID)
	Recommended string  `url:"recommended,omitempty"` //Filter with recommended source. If `CISCO` then the network device product assigned was recommended by Cisco and `USER` then the user has manually assigned. Available values: CISCO, USER
	Assigned    string  `url:"assigned,omitempty"`    //Filter with the assigned/unassigned, `ASSIGNED` option will filter network device products that are associated with the given image. The `NOT_ASSIGNED` option will filter network device products that have not yet been associated with the given image but apply to it. Available values: ASSIGNED, NOT_ASSIGNED
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1. The minimum value is 1
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively
}
type RetrievesTheCountOfAssignedNetworkDeviceProductsQueryParams struct {
	ProductName string `url:"productName,omitempty"` //Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search.
	ProductID   string `url:"productId,omitempty"`   //Filter with product ID (PID)
	Recommended string `url:"recommended,omitempty"` //Filter with recommended source. If `CISCO` then the network device product assigned was recommended by Cisco and `USER` then the user has manually assigned. Available values : CISCO, USER
	Assigned    string `url:"assigned,omitempty"`    //Filter with the assigned/unassigned, `ASSIGNED` option will filter network device products that are associated with the given image. The `NOT_ASSIGNED` option will filter network device products that have not yet been associated with the given image but apply to it. Available values: ASSIGNED, NOT_ASSIGNED
}
type GetNetworkDeviceImageUpdatesQueryParams struct {
	ID                string  `url:"id,omitempty"`                //Update id which is unique for each network device under the parentId
	ParentID          string  `url:"parentId,omitempty"`          //Updates that have this parent id
	NetworkDeviceID   string  `url:"networkDeviceId,omitempty"`   //Network device id
	Status            string  `url:"status,omitempty"`            //Status of the image update. Available values : FAILURE, SUCCESS, IN_PROGRESS, PENDING
	ImageName         string  `url:"imageName,omitempty"`         //Software image name for the update
	HostName          string  `url:"hostName,omitempty"`          //Host name of the network device for the image update. Supports case-insensitive partial search
	ManagementAddress string  `url:"managementAddress,omitempty"` //Management address of the network device
	StartTime         float64 `url:"startTime,omitempty"`         //Image update started after the given time (as milliseconds since UNIX epoch)
	EndTime           float64 `url:"endTime,omitempty"`           //Image update started before the given time (as milliseconds since UNIX epoch)
	SortBy            string  `url:"sortBy,omitempty"`            //A property within the response to sort by.
	Order             string  `url:"order,omitempty"`             //Whether ascending or descending order should be used to sort the response.
	Offset            float64 `url:"offset,omitempty"`            //The first record to show for this page; the first record is numbered 1.
	Limit             float64 `url:"limit,omitempty"`             //The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively
}
type CountOfNetworkDeviceImageUpdatesQueryParams struct {
	ID                string  `url:"id,omitempty"`                //Update id which is unique for each network device under the parentId
	ParentID          string  `url:"parentId,omitempty"`          //Updates that have this parent id
	NetworkDeviceID   string  `url:"networkDeviceId,omitempty"`   //Network device id
	Status            string  `url:"status,omitempty"`            //Status of the image update. Available values: FAILURE, SUCCESS, IN_PROGRESS, PENDING
	ImageName         string  `url:"imageName,omitempty"`         //Software image name for the update
	HostName          string  `url:"hostName,omitempty"`          //Host name of the network device for the image update. Supports case-insensitive partial search.
	ManagementAddress string  `url:"managementAddress,omitempty"` //Management address of the network device
	StartTime         float64 `url:"startTime,omitempty"`         //Image update started after the given time (as milliseconds since UNIX epoch).
	EndTime           float64 `url:"endTime,omitempty"`           //Image update started before the given time (as milliseconds since UNIX epoch).
}
type RetrievesTheListOfNetworkDeviceProductNamesQueryParams struct {
	ProductName string  `url:"productName,omitempty"` //Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search
	ProductID   string  `url:"productId,omitempty"`   //Filter with product ID (PID)
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1. The minimum value is 1.
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively.
}
type CountOfNetworkProductNamesQueryParams struct {
	ProductName string `url:"productName,omitempty"` //Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search
	ProductID   string `url:"productId,omitempty"`   //Filter with product ID (PID)
}
type ReturnsTheImageSummaryForTheGivenSiteQueryParams struct {
	SiteID string `url:"siteId,omitempty"` //Site identifier to get the aggreagte counts products under the site. The default value is global site id. See [https://developer.cisco.com/docs/dna-center](#!get-site) for `siteId`
}
type ReturnsNetworkDeviceProductNamesForASiteQueryParams struct {
	SiteID      string  `url:"siteId,omitempty"`      //Site identifier to get the list of all available products under the site. The default value is the global site.  See https://developer.cisco.com/docs/dna-center/get-site for siteId
	ProductName string  `url:"productName,omitempty"` //Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search
	Offset      float64 `url:"offset,omitempty"`      //The first record to show for this page; the first record is numbered 1. The minimum value is 1
	Limit       float64 `url:"limit,omitempty"`       //The number of records to show for this page. The minimum and maximum values are 1 and 500, respectively
}
type ReturnsTheCountOfNetworkDeviceProductNamesForASiteQueryParams struct {
	SiteID      string `url:"siteId,omitempty"`      //Site identifier to get the list of all available products under the site. The default value is global site id. See https://developer.cisco.com/docs/dna-center/get-site/ for siteId
	ProductName string `url:"productName,omitempty"` //Filter with network device product name. Supports partial case-insensitive search. A minimum of 3 characters are required for search
}

type ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivation struct {
	Response *ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivationResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivationResponse struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistribution struct {
	Response *ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionResponse struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetails struct {
	Response *[]ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponse struct {
	ApplicableDevicesForImage *[]ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseApplicableDevicesForImage `json:"applicableDevicesForImage,omitempty"` //

	ApplicationType string `json:"applicationType,omitempty"` //

	CreatedTime string `json:"createdTime,omitempty"` //

	ExtendedAttributes *ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseExtendedAttributes `json:"extendedAttributes,omitempty"` //

	Family string `json:"family,omitempty"` //

	Feature string `json:"feature,omitempty"` //

	FileServiceID string `json:"fileServiceId,omitempty"` //

	FileSize string `json:"fileSize,omitempty"` //

	ImageIntegrityStatus string `json:"imageIntegrityStatus,omitempty"` //

	ImageName string `json:"imageName,omitempty"` //

	ImageSeries []string `json:"imageSeries,omitempty"` //

	ImageSource string `json:"imageSource,omitempty"` //

	ImageType string `json:"imageType,omitempty"` //

	ImageUUID string `json:"imageUuid,omitempty"` //

	ImportSourceType string `json:"importSourceType,omitempty"` //

	IsTaggedGolden *bool `json:"isTaggedGolden,omitempty"` //

	Md5Checksum string `json:"md5Checksum,omitempty"` //

	Name string `json:"name,omitempty"` //

	ProfileInfo *[]ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseProfileInfo `json:"profileInfo,omitempty"` //

	ShaCheckSum string `json:"shaCheckSum,omitempty"` //

	Vendor string `json:"vendor,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseApplicableDevicesForImage struct {
	MdfID string `json:"mdfId,omitempty"` //

	ProductID []string `json:"productId,omitempty"` //

	ProductName string `json:"productName,omitempty"` //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseExtendedAttributes interface{}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseProfileInfo struct {
	Description string `json:"description,omitempty"` //

	ExtendedAttributes *ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseProfileInfoExtendedAttributes `json:"extendedAttributes,omitempty"` //

	Memory *int `json:"memory,omitempty"` //

	ProductType string `json:"productType,omitempty"` //

	ProfileName string `json:"profileName,omitempty"` //

	Shares *int `json:"shares,omitempty"` //

	VCPU *int `json:"vCpu,omitempty"` //
}
type ResponseSoftwareImageManagementSwimGetSoftwareImageDetailsResponseProfileInfoExtendedAttributes interface{}
type ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiers struct {
	Version string `json:"version,omitempty"` // Response Version

	Response *[]ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiersResponse `json:"response,omitempty"` //
}
type ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiersResponse struct {
	DeviceFamily string `json:"deviceFamily,omitempty"` // Device Family e.g. : Cisco Catalyst 6503 Switch-Catalyst 6500 Series Supervisor Engine 2T

	DeviceFamilyIDentifier string `json:"deviceFamilyIdentifier,omitempty"` // Device Family Identifier used for tagging an image golden for certain Device Family e.g. : 277696480-283933147
}
type ResponseSoftwareImageManagementSwimTagAsGoldenImage struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseSoftwareImageManagementSwimTagAsGoldenImageResponse `json:"response,omitempty"` //
}
type ResponseSoftwareImageManagementSwimTagAsGoldenImageResponse struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSoftwareImageManagementSwimRemoveGoldenTagForImage struct {
	Version string `json:"version,omitempty"` // Response Version e.g. : 1.0

	Response *ResponseSoftwareImageManagementSwimRemoveGoldenTagForImageResponse `json:"response,omitempty"` //
}
type ResponseSoftwareImageManagementSwimRemoveGoldenTagForImageResponse struct {
	URL string `json:"url,omitempty"` // URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5

	TaskID string `json:"taskId,omitempty"` // Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
}
type ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImage struct {
	Version string `json:"version,omitempty"` // Response Version. E.G. : 1.0

	Response *ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageResponse `json:"response,omitempty"` //
}
type ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImageResponse struct {
	DeviceRole string `json:"deviceRole,omitempty"` // Device Role. Possible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.

	TaggedGolden *bool `json:"taggedGolden,omitempty"` // Tagged Golden.

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // Inherited Site Name. If the Golden Tag is not tagged for the current site but is inherited from a higher enclosing site, it will contain the name of the site from where the tag is inherited.

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // Inherited Site Id. If the Golden Tag is not tagged for the current site but is inherited from a higher enclosing site, it will contain the uuid of the site from where the tag is inherited. In case the golden tag is inherited from the Global site the value will be "-1".
}
type ResponseSoftwareImageManagementSwimImportLocalSoftwareImage struct {
	Response *ResponseSoftwareImageManagementSwimImportLocalSoftwareImageResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseSoftwareImageManagementSwimImportLocalSoftwareImageResponse struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseSoftwareImageManagementSwimImportSoftwareImageViaURL struct {
	Response *ResponseSoftwareImageManagementSwimImportSoftwareImageViaURLResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` //
}
type ResponseSoftwareImageManagementSwimImportSoftwareImageViaURLResponse struct {
	TaskID string `json:"taskId,omitempty"` //

	URL string `json:"url,omitempty"` //
}
type ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImages struct {
	Response *[]ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesResponse struct {
	ID string `json:"id,omitempty"` // Software image identifier

	Imported *bool `json:"imported,omitempty"` // Flag for image info whether it is imported image or cloud image

	Name string `json:"name,omitempty"` // Name of the software image

	Version string `json:"version,omitempty"` // Software image  version

	ImageType string `json:"imageType,omitempty"` // Software image type

	Recommended string `json:"recommended,omitempty"` // CISCO if the image is recommended from Cisco.com

	CiscoLatest *bool `json:"ciscoLatest,omitempty"` // `true` if the image is latest/suggested from Cisco.com

	IntegrityStatus string `json:"integrityStatus,omitempty"` // Image Integrity verification status with Known Good Verification

	IsAddonImage *bool `json:"isAddonImage,omitempty"` // The value of `true` will indicate the image as an add-on image, while the value of `false` will indicate software image

	HasAddonImages *bool `json:"hasAddonImages,omitempty"` // Software images that have an applicable list of add-on images. The value of `true` will return software images with add-on images, while the value of `false` will return software images without add-on images

	GoldenTaggingDetails *[]ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesResponseGoldenTaggingDetails `json:"goldenTaggingDetails,omitempty"` //

	ProductNames *[]ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesResponseProductNames `json:"productNames,omitempty"` //

	IsGoldenTagged *bool `json:"isGoldenTagged,omitempty"` // The value of `true` will indicate the image marked as golden, while the value of `false` will indicate the image not marked as golden
}
type ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesResponseGoldenTaggingDetails struct {
	DeviceRoles []string `json:"deviceRoles,omitempty"` // Golden tagging based on the device roles

	DeviceTags []string `json:"deviceTags,omitempty"` // Golden tagging based on the device tags

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // The Site Id of the site that this setting is inherited from.

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // The name of the site that this setting is inherited from
}
type ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImagesResponseProductNames struct {
	ID string `json:"id,omitempty"` // Product name ordinal is unique value for each network device product

	ProductName string `json:"productName,omitempty"` // Network device product name

	ProductNameOrdinal *float64 `json:"productNameOrdinal,omitempty"` // Product name ordinal is unique value for each network device product

	SupervisorProductName string `json:"supervisorProductName,omitempty"` // Name of the Supervisor Engine Module, supported by the `productName`.                  Example: The `Cisco Catalyst 9404R Switch` chassis is capable of supporting  different supervisor engine modules: the `Cisco Catalyst 9400 Supervisor Engine-1`, the `Cisco Catalyst 9400 Supervisor Engine-1XL`, the `Cisco Catalyst 9400 Supervisor Engine-1XL-Y`, etc.

	SupervisorProductNameOrdinal *float64 `json:"supervisorProductNameOrdinal,omitempty"` // Supervisor Engine Module Ordinal, supported by the `productNameOrdinal`. Example: The `286315691` chassis ordinal is capable of supporting different supervisor engine module ordinals: `286316172`, `286316710`, `286320394` etc.
}
type ResponseSoftwareImageManagementSwimInitiatesSyncOfSoftwareImagesFromCiscoCom struct {
	Response *ResponseSoftwareImageManagementSwimInitiatesSyncOfSoftwareImagesFromCiscoComResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimInitiatesSyncOfSoftwareImagesFromCiscoComResponse struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task

	URL string `json:"url,omitempty"` // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImages struct {
	Response *ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImagesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImagesResponse struct {
	Count *int `json:"count,omitempty"` // Reports a count, for example, a total count of records in a given resource.
}
type ResponseSoftwareImageManagementSwimAddImageDistributionServer struct {
	Response *ResponseSoftwareImageManagementSwimAddImageDistributionServerResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimAddImageDistributionServerResponse struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task

	URL string `json:"url,omitempty"` // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimRetrieveImageDistributionServers struct {
	Response *[]ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimRetrieveImageDistributionServersResponse struct {
	ID string `json:"id,omitempty"` // Unique identifier for the server

	Username string `json:"username,omitempty"` // Server username

	ServerAddress string `json:"serverAddress,omitempty"` // FQDN or IP address of the server

	PortNumber *float64 `json:"portNumber,omitempty"` // Port number

	RootLocation string `json:"rootLocation,omitempty"` // Server root location
}
type ResponseSoftwareImageManagementSwimUpdateRemoteImageDistributionServer struct {
	Response *ResponseSoftwareImageManagementSwimUpdateRemoteImageDistributionServerResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimUpdateRemoteImageDistributionServerResponse struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task

	URL string `json:"url,omitempty"` // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimRetrieveSpecificImageDistributionServer struct {
	Response *ResponseSoftwareImageManagementSwimRetrieveSpecificImageDistributionServerResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimRetrieveSpecificImageDistributionServerResponse struct {
	ID string `json:"id,omitempty"` // Unique identifier for the server

	ServerAddress string `json:"serverAddress,omitempty"` // FQDN or IP address of the server

	PortNumber *float64 `json:"portNumber,omitempty"` // Port number

	RootLocation string `json:"rootLocation,omitempty"` // Server root location

	Username string `json:"username,omitempty"` // Server username
}
type ResponseSoftwareImageManagementSwimRemoveImageDistributionServer struct {
	Response *ResponseSoftwareImageManagementSwimRemoveImageDistributionServerResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimRemoveImageDistributionServerResponse struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task

	URL string `json:"url,omitempty"` // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimDeleteImage struct {
	Response *ResponseSoftwareImageManagementSwimDeleteImageResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimDeleteImageResponse struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task

	URL string `json:"url,omitempty"` // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImage struct {
	Response *[]ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version
}
type ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageResponse struct {
	ID string `json:"id,omitempty"` // Software image identifier

	Imported *bool `json:"imported,omitempty"` // Flag for image info whether it is imported image or cloud image

	Name string `json:"name,omitempty"` // Name of the software image

	Version string `json:"version,omitempty"` // Software image  version

	ImageType string `json:"imageType,omitempty"` // Software image type

	Recommended string `json:"recommended,omitempty"` // CISCO if the image is recommended from Cisco.com

	CiscoLatest *bool `json:"ciscoLatest,omitempty"` // `true` if the image is latest/suggested from Cisco.com

	IntegrityStatus string `json:"integrityStatus,omitempty"` // Image Integrity verification status with Known Good Verification

	IsAddonImage *bool `json:"isAddonImage,omitempty"` // The value of `true` will indicate the image as an add-on image, while the value of `false` will indicate software image

	HasAddonImages *bool `json:"hasAddonImages,omitempty"` // Software images that have an applicable list of add-on images. The value of `true` will return software images with add-on images, while the value of `false` will return software images without add-on images

	GoldenTaggingDetails *[]ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageResponseGoldenTaggingDetails `json:"goldenTaggingDetails,omitempty"` //

	ProductNames *[]ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageResponseProductNames `json:"productNames,omitempty"` //

	IsGoldenTagged *bool `json:"isGoldenTagged,omitempty"` // The value of `true` will indicate the image marked as golden, while the value of `false` will indicate the image not marked as golden
}
type ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageResponseGoldenTaggingDetails struct {
	DeviceRoles []string `json:"deviceRoles,omitempty"` // Golden tagging based on the device roles

	DeviceTags []string `json:"deviceTags,omitempty"` // Golden tagging based on the device tags

	InheritedSiteID string `json:"inheritedSiteId,omitempty"` // The Site Id of the site that this setting is inherited from.

	InheritedSiteName string `json:"inheritedSiteName,omitempty"` // The name of the site that this setting is inherited from
}
type ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImageResponseProductNames struct {
	ID string `json:"id,omitempty"` // Product name ordinal is unique value for each network device product

	ProductName string `json:"productName,omitempty"` // Network device product name

	ProductNameOrdinal *float64 `json:"productNameOrdinal,omitempty"` // Product name ordinal is unique value for each network device product

	SupervisorProductName string `json:"supervisorProductName,omitempty"` // Name of the Supervisor Engine Module, supported by the `productName`.                  Example: The `Cisco Catalyst 9404R Switch` chassis is capable of supporting  different supervisor engine modules: the `Cisco Catalyst 9400 Supervisor Engine-1`, the `Cisco Catalyst 9400 Supervisor Engine-1XL`, the `Cisco Catalyst 9400 Supervisor Engine-1XL-Y`, etc.

	SupervisorProductNameOrdinal *float64 `json:"supervisorProductNameOrdinal,omitempty"` // Supervisor Engine Module Ordinal, supported by the `productNameOrdinal`. Example: The `286315691` chassis ordinal is capable of supporting different supervisor engine module ordinals: `286316172`, `286316710`, `286320394` etc.
}
type ResponseSoftwareImageManagementSwimReturnsCountOfAddOnImages struct {
	Response *ResponseSoftwareImageManagementSwimReturnsCountOfAddOnImagesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseSoftwareImageManagementSwimReturnsCountOfAddOnImagesResponse struct {
	Count *int `json:"count,omitempty"` // Reports a count, for example, a total count of records in a given resource.
}
type ResponseSoftwareImageManagementSwimDownloadTheSoftwareImage struct {
	Response *ResponseSoftwareImageManagementSwimDownloadTheSoftwareImageResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimDownloadTheSoftwareImageResponse struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task

	URL string `json:"url,omitempty"` // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimTaggingGoldenImage struct {
	Response *ResponseSoftwareImageManagementSwimTaggingGoldenImageResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimTaggingGoldenImageResponse struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task

	URL string `json:"url,omitempty"` // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimUntaggingGoldenImage struct {
	Response *ResponseSoftwareImageManagementSwimUntaggingGoldenImageResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimUntaggingGoldenImageResponse struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task

	URL string `json:"url,omitempty"` // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImage struct {
	Response *ResponseSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImageResponse struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task

	URL string `json:"url,omitempty"` // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImage struct {
	Response *[]ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImageResponse struct {
	ID string `json:"id,omitempty"` // Product name ordinal is unique value for each network device product

	ProductName string `json:"productName,omitempty"` // Network device product name

	ProductNameOrdinal *float64 `json:"productNameOrdinal,omitempty"` // Product name ordinal is unique value for each network device product

	ProductIDs []string `json:"productIds,omitempty"` // Supported PIDs

	SiteIDs []string `json:"siteIds,omitempty"` // Sites where all  this image is assigned

	Recommended string `json:"recommended,omitempty"` // If 'CISCO' network device product recommandation came from Cisco.com and 'USER' manually assigned
}
type ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProducts struct {
	Response *ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProductsResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProductsResponse struct {
	Count *int `json:"count,omitempty"` // Reports a count, for example, a total count of records in a given resource.
}
type ResponseSoftwareImageManagementSwimUnassignNetworkDeviceProductNameFromTheGivenSoftwareImage struct {
	Response *ResponseSoftwareImageManagementSwimUnassignNetworkDeviceProductNameFromTheGivenSoftwareImageResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimUnassignNetworkDeviceProductNameFromTheGivenSoftwareImageResponse struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task

	URL string `json:"url,omitempty"` // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage struct {
	Response *ResponseSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImageResponse struct {
	TaskID string `json:"taskId,omitempty"` // The UUID of the task

	URL string `json:"url,omitempty"` // The path to the API endpoint to GET for information on the task
}
type ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdates struct {
	Response *[]ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdatesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // API response version
}
type ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdatesResponse struct {
	ID string `json:"id,omitempty"` // Unique identifier for the image update

	ParentID string `json:"parentId,omitempty"` // Parent identifier for the image update

	StartTime *float64 `json:"startTime,omitempty"` // Image update started after the given time (as milliseconds since UNIX epoch)

	EndTime *float64 `json:"endTime,omitempty"` // Image update end time (as milliseconds since UNIX epoch)

	Status string `json:"status,omitempty"` // Status of the image update

	NetworkDeviceID string `json:"networkDeviceId,omitempty"` // Network device identifier

	ManagementAddress string `json:"managementAddress,omitempty"` // Management address of the network device

	HostName string `json:"hostName,omitempty"` // Host name of the network device for the image update

	UpdateImageVersion string `json:"updateImageVersion,omitempty"` // Software image version

	Type string `json:"type,omitempty"` // Type of the image update
}
type ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdates struct {
	Response *ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdatesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdatesResponse struct {
	Count *int `json:"count,omitempty"` // Reports a count, for example, a total count of records in a given resource.
}
type ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNames struct {
	Response *[]ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNamesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNamesResponse struct {
	ID string `json:"id,omitempty"` // Product name ordinal is unique value for each network device product

	ProductName string `json:"productName,omitempty"` // Network device product name

	ProductNameOrdinal *float64 `json:"productNameOrdinal,omitempty"` // Product name ordinal is unique value for each network device product

	ProductIDs []string `json:"productIds,omitempty"` // Supported PIDs
}
type ResponseSoftwareImageManagementSwimCountOfNetworkProductNames struct {
	Response *ResponseSoftwareImageManagementSwimCountOfNetworkProductNamesResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseSoftwareImageManagementSwimCountOfNetworkProductNamesResponse struct {
	Count *int `json:"count,omitempty"` // Reports a count, for example, a total count of records in a given resource.
}
type ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductName struct {
	Response *ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductNameResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductNameResponse struct {
	ID string `json:"id,omitempty"` // Product name ordinal is unique value for each network device product

	ProductName string `json:"productName,omitempty"` // Network device product name

	ProductNameOrdinal *float64 `json:"productNameOrdinal,omitempty"` // Product name ordinal is unique value for each network device product

	ProductIDs []string `json:"productIds,omitempty"` // Supported PIDs
}
type ResponseSoftwareImageManagementSwimReturnsTheImageSummaryForTheGivenSite struct {
	Response *ResponseSoftwareImageManagementSwimReturnsTheImageSummaryForTheGivenSiteResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Version of the response
}
type ResponseSoftwareImageManagementSwimReturnsTheImageSummaryForTheGivenSiteResponse struct {
	ImportedImageCount *float64 `json:"importedImageCount,omitempty"` // Count of images imported

	InstalledImageCount *float64 `json:"installedImageCount,omitempty"` // Count of installed images

	GoldenImageCount *float64 `json:"goldenImageCount,omitempty"` // Count of images marked as golden

	NonGoldenImageCount *float64 `json:"nonGoldenImageCount,omitempty"` // Count of images not marked as golden

	InstalledImageAdvisorCount *float64 `json:"installedImageAdvisorCount,omitempty"` // Advisor count of installed images

	ProductCount *float64 `json:"productCount,omitempty"` // Count of Network device product name

	ProductsWithGoldenCount *float64 `json:"productsWithGoldenCount,omitempty"` // Count of Network device product name marked as golden

	ProductsWithoutGoldenCount *float64 `json:"productsWithoutGoldenCount,omitempty"` // Count of Network device product name not marked as golden
}
type ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASite struct {
	Response *ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // Response version
}
type ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteResponse struct {
	ID string `json:"id,omitempty"` // The unique identifier for the record is the `id`. If there is no supervisor engine involved, the `id` will be the same as the `productNameOrdinal`. However, if the supervisor engine is applicable, the `id` will be in the form of `<productNameOrdinal>-<supervisorProductNameOrdinal>`.

	ProductNameOrdinal *float64 `json:"productNameOrdinal,omitempty"` // Product name ordinal

	ProductName string `json:"productName,omitempty"` // Name of product

	SupervisorProductName string `json:"supervisorProductName,omitempty"` // Name of the Supervisor Engine Module, supported by the `productName`.                  Example: The `Cisco Catalyst 9404R Switch` chassis is capable of supporting  different supervisor engine modules: the `Cisco Catalyst 9400 Supervisor Engine-1`, the `Cisco Catalyst 9400 Supervisor Engine-1XL`, the `Cisco Catalyst 9400 Supervisor Engine-1XL-Y`, etc.

	SupervisorProductNameOrdinal *float64 `json:"supervisorProductNameOrdinal,omitempty"` // Supervisor Engine Module Ordinal, supported by the `productNameOrdinal`. Example: The `286315691` chassis ordinal is capable of supporting  different supervisor engine module ordinals: `286316172`, `286316710`, `286320394` etc.

	NetworkDeviceCount *int `json:"networkDeviceCount,omitempty"` // Count of network devices

	ImageSummary *ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteResponseImageSummary `json:"imageSummary,omitempty"` //
}
type ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASiteResponseImageSummary struct {
	InstalledImageCount *int `json:"installedImageCount,omitempty"` // Count of installed images

	GoldenImageCount *int `json:"goldenImageCount,omitempty"` // Count of golden tagged images

	InstalledImageAdvisorCount *int `json:"installedImageAdvisorCount,omitempty"` // Count of advisor on installed images
}
type ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASite struct {
	Response *ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASiteResponse `json:"response,omitempty"` //

	Version string `json:"version,omitempty"` // The version of the response
}
type ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASiteResponse struct {
	Count *int `json:"count,omitempty"` // Reports a count, for example, a total count of records in a given resource.
}
type RequestSoftwareImageManagementSwimTriggerSoftwareImageActivation []RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivation // Array of RequestSoftwareImageManagementSwimTriggerSoftwareImageActivation
type RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivation struct {
	ActivateLowerImageVersion *bool `json:"activateLowerImageVersion,omitempty"` //

	DeviceUpgradeMode string `json:"deviceUpgradeMode,omitempty"` //

	DeviceUUID string `json:"deviceUuid,omitempty"` //

	DistributeIfNeeded *bool `json:"distributeIfNeeded,omitempty"` //

	ImageUUIDList []string `json:"imageUuidList,omitempty"` //

	SmuImageUUIDList []string `json:"smuImageUuidList,omitempty"` //
}
type RequestSoftwareImageManagementSwimTriggerSoftwareImageDistribution []RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistribution // Array of RequestSoftwareImageManagementSwimTriggerSoftwareImageDistribution
type RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistribution struct {
	DeviceUUID string `json:"deviceUuid,omitempty"` //

	ImageUUID string `json:"imageUuid,omitempty"` //
}
type RequestSoftwareImageManagementSwimTagAsGoldenImage struct {
	ImageID string `json:"imageId,omitempty"` // imageId in uuid format.

	SiteID string `json:"siteId,omitempty"` // SiteId in uuid format. For Global Site "-1" to be used.

	DeviceRole string `json:"deviceRole,omitempty"` // Device Role. Permissible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.

	DeviceFamilyIDentifier string `json:"deviceFamilyIdentifier,omitempty"` // Device Family Identifier e.g. : 277696480-283933147, 277696480
}
type RequestSoftwareImageManagementSwimImportSoftwareImageViaURL []RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURL // Array of RequestSoftwareImageManagementSwimImportSoftwareImageViaURL
type RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURL struct {
	ApplicationType string `json:"applicationType,omitempty"` //

	ImageFamily string `json:"imageFamily,omitempty"` //

	SourceURL string `json:"sourceURL,omitempty"` //

	ThirdParty *bool `json:"thirdParty,omitempty"` //

	Vendor string `json:"vendor,omitempty"` //
}
type RequestSoftwareImageManagementSwimAddImageDistributionServer struct {
	ServerAddress string `json:"serverAddress,omitempty"` // FQDN or IP address of the server

	Username string `json:"username,omitempty"` // Server username

	PortNumber *float64 `json:"portNumber,omitempty"` // Port number

	RootLocation string `json:"rootLocation,omitempty"` // Server root location

	Password string `json:"password,omitempty"` // Server password
}
type RequestSoftwareImageManagementSwimUpdateRemoteImageDistributionServer struct {
	Username string `json:"username,omitempty"` // Server username

	PortNumber *float64 `json:"portNumber,omitempty"` // Port number

	Password string `json:"password,omitempty"` // Server password
}
type RequestSoftwareImageManagementSwimTaggingGoldenImage struct {
	ProductNameOrdinal *float64 `json:"productNameOrdinal,omitempty"` // The product name ordinal is a unique value for each network device product. `productNameOrdinal` can be obtained from the response of API `/dna/intent/api/v1/siteWiseProductNames?siteId=<siteId>`

	SupervisorProductNameOrdinal *float64 `json:"supervisorProductNameOrdinal,omitempty"` // The supervisor engine module ordinal is a unique value for each supervisor module. `supervisorProductNameOrdinal` can be obtained from the response of API `/dna/intent/api/v1/siteWiseProductNames?siteId=<siteId>`

	DeviceRoles []string `json:"deviceRoles,omitempty"` // Device Roles. Available value will be [ CORE, DISTRIBUTION, UNKNOWN, ACCESS, BORDER ROUTER ]

	DeviceTags []string `json:"deviceTags,omitempty"` // Device tags can be fetched fom API https://developer.cisco.com/docs/dna-center/#!get-tag
}
type RequestSoftwareImageManagementSwimUntaggingGoldenImage struct {
	ProductNameOrdinal *float64 `json:"productNameOrdinal,omitempty"` // The product name ordinal is a unique value for each network device product. `productNameOrdinal` can be obtained from the response of API `/dna/intent/api/v1/siteWiseProductNames?siteId=<siteId>`

	SupervisorProductNameOrdinal *float64 `json:"supervisorProductNameOrdinal,omitempty"` // The supervisor engine module ordinal is a unique value for each supervisor module. `supervisorProductNameOrdinal` can be obtained from the response of API `/dna/intent/api/v1/siteWiseProductNames?siteId=<siteId>`

	DeviceRoles []string `json:"deviceRoles,omitempty"` // Device Roles. Available value will be [ CORE, DISTRIBUTION, UNKNOWN, ACCESS, BORDER ROUTER ]

	DeviceTags []string `json:"deviceTags,omitempty"` // Device tags can be fetched fom API https://developer.cisco.com/docs/dna-center/#!get-tag
}
type RequestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImage struct {
	ProductNameOrdinal *float64 `json:"productNameOrdinal,omitempty"` // Product name ordinal is unique value for each network device product

	SiteIDs []string `json:"siteIds,omitempty"` // Sites where this image needs to be assigned. Ref https://developer.cisco.com/docs/dna-center/#!sites
}
type RequestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage struct {
	SiteIDs []string `json:"siteIds,omitempty"` // Sites where all this image need to be assigned
}

//GetSoftwareImageDetails Get software image details - 0c8f-7a0b-49b9-aedd
/* Returns software image list based on a filter criteria. For example: "filterbyName = cat3k%"


@param GetSoftwareImageDetailsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-software-image-details
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetSoftwareImageDetails(GetSoftwareImageDetailsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetSoftwareImageDetails")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimGetSoftwareImageDetails)
	return result, response, err

}

//GetDeviceFamilyIDentifiers Get Device Family Identifiers - 35ae-1bec-4bd8-89fc
/* API to get Device Family Identifiers for all Device Families that can be used for tagging an image golden.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-device-family-identifiers
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetDeviceFamilyIDentifiers()
		}
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


Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-golden-tag-status-of-an-image
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetGoldenTagStatusOfAnImage(siteID, deviceFamilyIDentifier, deviceRole, imageID)
		}
		return nil, response, fmt.Errorf("error with operation GetGoldenTagStatusOfAnImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimGetGoldenTagStatusOfAnImage)
	return result, response, err

}

//ReturnsListOfSoftwareImages Returns list of software images - e4a3-6a8c-48fa-91ea
/* A list of available images for the specified site is provided. The default value of the site is set to global. The list includes images that have been imported onto the disk, as well as the latest and suggested images from Cisco.com.


@param ReturnsListOfSoftwareImagesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-list-of-software-images
*/
func (s *SoftwareImageManagementSwimService) ReturnsListOfSoftwareImages(ReturnsListOfSoftwareImagesQueryParams *ReturnsListOfSoftwareImagesQueryParams) (*ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImages, *resty.Response, error) {
	path := "/dna/intent/api/v1/images"

	queryString, _ := query.Values(ReturnsListOfSoftwareImagesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImages{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsListOfSoftwareImages(ReturnsListOfSoftwareImagesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsListOfSoftwareImages")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimReturnsListOfSoftwareImages)
	return result, response, err

}

//ReturnsCountOfSoftwareImages Returns count of software images - 1391-aa45-4098-8eac
/* Returns the count of software images for given `siteId`. The default value of siteId is global


@param ReturnsCountOfSoftwareImagesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-count-of-software-images
*/
func (s *SoftwareImageManagementSwimService) ReturnsCountOfSoftwareImages(ReturnsCountOfSoftwareImagesQueryParams *ReturnsCountOfSoftwareImagesQueryParams) (*ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImages, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/count"

	queryString, _ := query.Values(ReturnsCountOfSoftwareImagesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImages{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsCountOfSoftwareImages(ReturnsCountOfSoftwareImagesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsCountOfSoftwareImages")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimReturnsCountOfSoftwareImages)
	return result, response, err

}

//RetrieveImageDistributionServers Retrieve image distribution servers - 7982-39ee-4aaa-aa72
/* Retrieve the list of remote image distribution servers. There can be up to two remote servers.Product always acts as local distribution server, and it is not part of this API response.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-image-distribution-servers
*/
func (s *SoftwareImageManagementSwimService) RetrieveImageDistributionServers() (*ResponseSoftwareImageManagementSwimRetrieveImageDistributionServers, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/distributionServerSettings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimRetrieveImageDistributionServers{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveImageDistributionServers()
		}
		return nil, response, fmt.Errorf("error with operation RetrieveImageDistributionServers")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrieveImageDistributionServers)
	return result, response, err

}

//RetrieveSpecificImageDistributionServer Retrieve specific image distribution server - b1ac-99fe-47a9-9c85
/* Retrieve image distribution server for the given server identifier


@param id id path parameter. Server identifier


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-specific-image-distribution-server
*/
func (s *SoftwareImageManagementSwimService) RetrieveSpecificImageDistributionServer(id string) (*ResponseSoftwareImageManagementSwimRetrieveSpecificImageDistributionServer, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/distributionServerSettings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimRetrieveSpecificImageDistributionServer{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveSpecificImageDistributionServer(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveSpecificImageDistributionServer")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrieveSpecificImageDistributionServer)
	return result, response, err

}

//RetrieveApplicableAddOnImagesForTheGivenSoftwareImage Retrieve applicable add-on images for the given software image - f0ae-9b76-4ee8-ac84
/* Retrieves the list of applicable add-on images if available for the given software image. `id` can be obtained from the response of API [ /dna/intent/api/v1/images?hasAddonImages=true ].


@param id id path parameter. Software image identifier. Check `/dna/intent/api/v1/images?hasAddonImages=true` API to get the same.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-applicable-add-on-images-for-the-given-software-image
*/
func (s *SoftwareImageManagementSwimService) RetrieveApplicableAddOnImagesForTheGivenSoftwareImage(id string) (*ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImage, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{id}/addonImages"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImage{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveApplicableAddOnImagesForTheGivenSoftwareImage(id)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveApplicableAddOnImagesForTheGivenSoftwareImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrieveApplicableAddOnImagesForTheGivenSoftwareImage)
	return result, response, err

}

//ReturnsCountOfAddOnImages Returns count of add-on images - f0ba-68e0-4acb-8234
/* Count of add-on images available for the given software image identifier, `id` can be obtained from the response of API [ /dna/intent/api/v1/images?hasAddonImages=true ].


@param id id path parameter. Software image identifier. Check API `/dna/intent/api/v1/images` for id from response.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-count-of-add-on-images
*/
func (s *SoftwareImageManagementSwimService) ReturnsCountOfAddOnImages(id string) (*ResponseSoftwareImageManagementSwimReturnsCountOfAddOnImages, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{id}/addonImages/count"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimReturnsCountOfAddOnImages{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsCountOfAddOnImages(id)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsCountOfAddOnImages")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimReturnsCountOfAddOnImages)
	return result, response, err

}

//RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage Retrieves network device product names assigned to a software image. - 14a1-f9d8-49b9-80be
/* Returns a list of network device product names and associated sites for a given image identifier. Refer `/dna/intent/api/v1/images` API for obtaining `imageId`.


@param imageID imageId path parameter. Software image identifier. Refer `/dna/intent/api/v1/images` API for obtaining `imageId`

@param RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-network-device-product-names-assigned-to-a-software-image
*/
func (s *SoftwareImageManagementSwimService) RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage(imageID string, RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageQueryParams *RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageQueryParams) (*ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImage, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{imageId}/siteWiseProductNames"
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)

	queryString, _ := query.Values(RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImage{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage(imageID, RetrievesNetworkDeviceProductNamesAssignedToASoftwareImageQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesNetworkDeviceProductNamesAssignedToASoftwareImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrievesNetworkDeviceProductNamesAssignedToASoftwareImage)
	return result, response, err

}

//RetrievesTheCountOfAssignedNetworkDeviceProducts Retrieves the count of assigned network device products - e994-0a33-409b-b8be
/* Returns count of assigned network device product for a given image identifier. Refer `/dna/intent/api/v1/images` API for obtaining `imageId`


@param imageID imageId path parameter. Software image identifier. Refer `/dna/intent/api/v/images` API for obtaining `imageId`

@param RetrievesTheCountOfAssignedNetworkDeviceProductsQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-count-of-assigned-network-device-products
*/
func (s *SoftwareImageManagementSwimService) RetrievesTheCountOfAssignedNetworkDeviceProducts(imageID string, RetrievesTheCountOfAssignedNetworkDeviceProductsQueryParams *RetrievesTheCountOfAssignedNetworkDeviceProductsQueryParams) (*ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProducts, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{imageId}/siteWiseProductNames/count"
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)

	queryString, _ := query.Values(RetrievesTheCountOfAssignedNetworkDeviceProductsQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProducts{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheCountOfAssignedNetworkDeviceProducts(imageID, RetrievesTheCountOfAssignedNetworkDeviceProductsQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheCountOfAssignedNetworkDeviceProducts")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrievesTheCountOfAssignedNetworkDeviceProducts)
	return result, response, err

}

//GetNetworkDeviceImageUpdates Get network device image updates - 75bf-3a9e-467b-af24
/* Returns the list of network device image updates based on the given filter criteria


@param GetNetworkDeviceImageUpdatesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!get-network-device-image-updates
*/
func (s *SoftwareImageManagementSwimService) GetNetworkDeviceImageUpdates(GetNetworkDeviceImageUpdatesQueryParams *GetNetworkDeviceImageUpdatesQueryParams) (*ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdates, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceImageUpdates"

	queryString, _ := query.Values(GetNetworkDeviceImageUpdatesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdates{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.GetNetworkDeviceImageUpdates(GetNetworkDeviceImageUpdatesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation GetNetworkDeviceImageUpdates")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimGetNetworkDeviceImageUpdates)
	return result, response, err

}

//CountOfNetworkDeviceImageUpdates Count of network device image updates - b980-b848-45a8-9987
/* Returns the count of network device image updates based on the given filter criteria


@param CountOfNetworkDeviceImageUpdatesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-of-network-device-image-updates
*/
func (s *SoftwareImageManagementSwimService) CountOfNetworkDeviceImageUpdates(CountOfNetworkDeviceImageUpdatesQueryParams *CountOfNetworkDeviceImageUpdatesQueryParams) (*ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdates, *resty.Response, error) {
	path := "/dna/intent/api/v1/networkDeviceImageUpdates/count"

	queryString, _ := query.Values(CountOfNetworkDeviceImageUpdatesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdates{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountOfNetworkDeviceImageUpdates(CountOfNetworkDeviceImageUpdatesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountOfNetworkDeviceImageUpdates")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimCountOfNetworkDeviceImageUpdates)
	return result, response, err

}

//RetrievesTheListOfNetworkDeviceProductNames Retrieves the list of network device product names - a7bf-3baf-4c29-b1ca
/* Get the list of network device product names, their ordinal, and the support PIDs based on filter criteria.


@param RetrievesTheListOfNetworkDeviceProductNamesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieves-the-list-of-network-device-product-names
*/
func (s *SoftwareImageManagementSwimService) RetrievesTheListOfNetworkDeviceProductNames(RetrievesTheListOfNetworkDeviceProductNamesQueryParams *RetrievesTheListOfNetworkDeviceProductNamesQueryParams) (*ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNames, *resty.Response, error) {
	path := "/dna/intent/api/v1/productNames"

	queryString, _ := query.Values(RetrievesTheListOfNetworkDeviceProductNamesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNames{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrievesTheListOfNetworkDeviceProductNames(RetrievesTheListOfNetworkDeviceProductNamesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation RetrievesTheListOfNetworkDeviceProductNames")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrievesTheListOfNetworkDeviceProductNames)
	return result, response, err

}

//CountOfNetworkProductNames Count of network product names - baa2-9b3d-45bb-a870
/* Count of product names based on filter criteria


@param CountOfNetworkProductNamesQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!count-of-network-product-names
*/
func (s *SoftwareImageManagementSwimService) CountOfNetworkProductNames(CountOfNetworkProductNamesQueryParams *CountOfNetworkProductNamesQueryParams) (*ResponseSoftwareImageManagementSwimCountOfNetworkProductNames, *resty.Response, error) {
	path := "/dna/intent/api/v1/productNames/count"

	queryString, _ := query.Values(CountOfNetworkProductNamesQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimCountOfNetworkProductNames{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.CountOfNetworkProductNames(CountOfNetworkProductNamesQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation CountOfNetworkProductNames")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimCountOfNetworkProductNames)
	return result, response, err

}

//RetrieveNetworkDeviceProductName Retrieve network device product name - 3aa8-fb90-4288-b606
/* Get the network device product name, its ordinal, and supported PIDs.


@param productNameOrdinal productNameOrdinal path parameter. Product name ordinal is unique value for each network device product.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!retrieve-network-device-product-name
*/
func (s *SoftwareImageManagementSwimService) RetrieveNetworkDeviceProductName(productNameOrdinal float64) (*ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductName, *resty.Response, error) {
	path := "/dna/intent/api/v1/productNames/{productNameOrdinal}"
	path = strings.Replace(path, "{productNameOrdinal}", fmt.Sprintf("%v", productNameOrdinal), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductName{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RetrieveNetworkDeviceProductName(productNameOrdinal)
		}
		return nil, response, fmt.Errorf("error with operation RetrieveNetworkDeviceProductName")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRetrieveNetworkDeviceProductName)
	return result, response, err

}

//ReturnsTheImageSummaryForTheGivenSite Returns the image summary for the given site - 6e96-e84e-43fb-baa8
/* Returns aggregate counts of network device product names, golden and non-golden tagged products, imported images, golden images tagged, and advisor for a specific site provide, the default value of `siteId` is set to global.


@param ReturnsTheImageSummaryForTheGivenSiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-the-image-summary-for-the-given-site
*/
func (s *SoftwareImageManagementSwimService) ReturnsTheImageSummaryForTheGivenSite(ReturnsTheImageSummaryForTheGivenSiteQueryParams *ReturnsTheImageSummaryForTheGivenSiteQueryParams) (*ResponseSoftwareImageManagementSwimReturnsTheImageSummaryForTheGivenSite, *resty.Response, error) {
	path := "/dna/intent/api/v1/siteWiseImagesSummary"

	queryString, _ := query.Values(ReturnsTheImageSummaryForTheGivenSiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimReturnsTheImageSummaryForTheGivenSite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsTheImageSummaryForTheGivenSite(ReturnsTheImageSummaryForTheGivenSiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsTheImageSummaryForTheGivenSite")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimReturnsTheImageSummaryForTheGivenSite)
	return result, response, err

}

//ReturnsNetworkDeviceProductNamesForASite Returns network device product names for a site - 20b5-5b0c-4518-9a03
/* Provides network device product names for a site. The default value of `siteId` is global. The response will include the network device count and image summary.


@param ReturnsNetworkDeviceProductNamesForASiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-network-device-product-names-for-a-site
*/
func (s *SoftwareImageManagementSwimService) ReturnsNetworkDeviceProductNamesForASite(ReturnsNetworkDeviceProductNamesForASiteQueryParams *ReturnsNetworkDeviceProductNamesForASiteQueryParams) (*ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/siteWiseProductNames"

	queryString, _ := query.Values(ReturnsNetworkDeviceProductNamesForASiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsNetworkDeviceProductNamesForASite(ReturnsNetworkDeviceProductNamesForASiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsNetworkDeviceProductNamesForASite")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimReturnsNetworkDeviceProductNamesForASite)
	return result, response, err

}

//ReturnsTheCountOfNetworkDeviceProductNamesForASite Returns the count of network device product names for a site - 018d-a93c-4e4b-8436
/* Returns the count of network device product names for given filters. The default value of `siteId` is global.


@param ReturnsTheCountOfNetworkDeviceProductNamesForASiteQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!returns-the-count-of-network-device-product-names-for-a-site
*/
func (s *SoftwareImageManagementSwimService) ReturnsTheCountOfNetworkDeviceProductNamesForASite(ReturnsTheCountOfNetworkDeviceProductNamesForASiteQueryParams *ReturnsTheCountOfNetworkDeviceProductNamesForASiteQueryParams) (*ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASite, *resty.Response, error) {
	path := "/dna/intent/api/v1/siteWiseProductNames/count"

	queryString, _ := query.Values(ReturnsTheCountOfNetworkDeviceProductNamesForASiteQueryParams)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetQueryString(queryString.Encode()).SetResult(&ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASite{}).
		SetError(&Error).
		Get(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.ReturnsTheCountOfNetworkDeviceProductNamesForASite(ReturnsTheCountOfNetworkDeviceProductNamesForASiteQueryParams)
		}
		return nil, response, fmt.Errorf("error with operation ReturnsTheCountOfNetworkDeviceProductNamesForASite")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimReturnsTheCountOfNetworkDeviceProductNamesForASite)
	return result, response, err

}

//TriggerSoftwareImageActivation Trigger software image activation - fb9b-eb66-4f2a-ba4c
/* Activates a software image on a given device. Software image must be present in the device flash


@param TriggerSoftwareImageActivationHeaderParams Custom header parameters
@param TriggerSoftwareImageActivationQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!trigger-software-image-activation
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

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TriggerSoftwareImageActivation(requestSoftwareImageManagementSwimTriggerSoftwareImageActivation, TriggerSoftwareImageActivationHeaderParams, TriggerSoftwareImageActivationQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation TriggerSoftwareImageActivation")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivation)
	return result, response, err

}

//TriggerSoftwareImageDistribution Trigger software image distribution - 8cb6-783b-4fab-a1f4
/* Distributes a software image on a given device. Software image must be imported successfully into DNA Center before it can be distributed



Documentation Link: https://developer.cisco.com/docs/dna-center/#!trigger-software-image-distribution
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

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TriggerSoftwareImageDistribution(requestSoftwareImageManagementSwimTriggerSoftwareImageDistribution)
		}

		return nil, response, fmt.Errorf("error with operation TriggerSoftwareImageDistribution")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistribution)
	return result, response, err

}

//TagAsGoldenImage Tag as Golden Image - 5c91-7a67-474b-a0e0
/* Golden Tag image. Set siteId as -1 for Global site.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!tag-as-golden-image
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

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TagAsGoldenImage(requestSoftwareImageManagementSwimTagAsGoldenImage)
		}

		return nil, response, fmt.Errorf("error with operation TagAsGoldenImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimTagAsGoldenImage)
	return result, response, err

}

//ImportLocalSoftwareImage Import local software image - 1491-f90f-48da-aabe
/* Fetches a software image from local file system and uploads to DNA Center. Supported software image files extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2


@param ImportLocalSoftwareImageQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-local-software-image
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

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportLocalSoftwareImage(ImportLocalSoftwareImageQueryParams, ImportLocalSoftwareImageMultipartFields)
		}

		return nil, response, fmt.Errorf("error with operation ImportLocalSoftwareImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimImportLocalSoftwareImage)
	return result, response, err

}

//ImportSoftwareImageViaURL Import software image via URL - bc8a-ab47-46ca-883d
/* Fetches a software image from remote file system (using URL for HTTP/FTP) and uploads to DNA Center. Supported image files extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2


@param ImportSoftwareImageViaURLQueryParams Filtering parameter

Documentation Link: https://developer.cisco.com/docs/dna-center/#!import-software-image-via-url
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

		if response.StatusCode() == http.StatusUnauthorized {
			return s.ImportSoftwareImageViaURL(requestSoftwareImageManagementSwimImportSoftwareImageViaURL, ImportSoftwareImageViaURLQueryParams)
		}

		return nil, response, fmt.Errorf("error with operation ImportSoftwareImageViaUrl")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimImportSoftwareImageViaURL)
	return result, response, err

}

//InitiatesSyncOfSoftwareImagesFromCiscoCom Initiates sync of software images from Cisco.com - 2b86-48e6-4b2a-b6f6
/* Initiating the synchronization of the software images from Cisco.com. The latest and suggested images will be retrieved, along with the corresponding product name and PIDs for imported and retrieved images from Cisco.com. Once the task is completed, the API `/intent/api/v1/images?imported=false` will display all the images fetched from Cisco.com.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!initiates-sync-of-software-images-from-cisco-com
*/
func (s *SoftwareImageManagementSwimService) InitiatesSyncOfSoftwareImagesFromCiscoCom() (*ResponseSoftwareImageManagementSwimInitiatesSyncOfSoftwareImagesFromCiscoCom, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/ccoSync"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimInitiatesSyncOfSoftwareImagesFromCiscoCom{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.InitiatesSyncOfSoftwareImagesFromCiscoCom()
		}

		return nil, response, fmt.Errorf("error with operation InitiatesSyncOfSoftwareImagesFromCiscoCom")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimInitiatesSyncOfSoftwareImagesFromCiscoCom)
	return result, response, err

}

//AddImageDistributionServer Add image distribution server - 0699-0a83-4aaa-be15
/* Add remote server for distributing software images. Upto two such distribution servers are supported.



Documentation Link: https://developer.cisco.com/docs/dna-center/#!add-image-distribution-server
*/
func (s *SoftwareImageManagementSwimService) AddImageDistributionServer(requestSoftwareImageManagementSwimAddImageDistributionServer *RequestSoftwareImageManagementSwimAddImageDistributionServer) (*ResponseSoftwareImageManagementSwimAddImageDistributionServer, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/distributionServerSettings"

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimAddImageDistributionServer).
		SetResult(&ResponseSoftwareImageManagementSwimAddImageDistributionServer{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AddImageDistributionServer(requestSoftwareImageManagementSwimAddImageDistributionServer)
		}

		return nil, response, fmt.Errorf("error with operation AddImageDistributionServer")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimAddImageDistributionServer)
	return result, response, err

}

//DownloadTheSoftwareImage Download the software image - c382-2b30-447a-a189
/* Initiates download of the software image from Cisco.com on the disk for the given `id`. Refer to `/dna/intent/api/v1/images` for obtaining `id`.


@param id id path parameter. Software image identifier. Check API `/dna/intent/api/v1/images` for `id` from response.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!download-the-software-image
*/
func (s *SoftwareImageManagementSwimService) DownloadTheSoftwareImage(id string) (*ResponseSoftwareImageManagementSwimDownloadTheSoftwareImage, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{id}/download"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimDownloadTheSoftwareImage{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.DownloadTheSoftwareImage(id)
		}

		return nil, response, fmt.Errorf("error with operation DownloadTheSoftwareImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimDownloadTheSoftwareImage)
	return result, response, err

}

//TaggingGoldenImage Tagging golden image - 24b1-da62-4c69-9a8c
/* Creates golden image tagging specifically for a particular device type or supervisor engine module. Conditions for tagging the golden image:
1) The golden tagging of SMU, PISRT_SMU, APSP, and APDP image type depends on the golden tagged applied on the base image. If any discrepancies are identified in the request payload, the golden tagging process will fail. For example:

    a) If the base image is tagged with Device Role: ACCESS, then add-ons can only be done ACCESS role only and the same is applied if any device tag is used. Any other request will fail.

    b) If the base image is tagged at global or any site level then add-ons also need to be tagged at site level.

2) Tagging of SUBPACKAGE and ROMMON image type is not supported.


@param id id path parameter. Software image identifier is used for golden tagging or intent to tag it. The value of `id` can be obtained from the response of the API `/dna/intent/api/v1/images?imported=true&isAddonImages=false` for the base image and `/dna/images/{id}/addonImages` where `id` will be the software image identifier of the base image.

@param siteID siteId path parameter. Site identifier for tagged image or intent to tag it. The default value is global site id. See [https://developer.cisco.com/docs/dna-center](#!get-site) for `siteId`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!tagging-golden-image
*/
func (s *SoftwareImageManagementSwimService) TaggingGoldenImage(id string, siteID string, requestSoftwareImageManagementSwimTaggingGoldenImage *RequestSoftwareImageManagementSwimTaggingGoldenImage) (*ResponseSoftwareImageManagementSwimTaggingGoldenImage, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{id}/sites/{siteId}/tagGolden"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimTaggingGoldenImage).
		SetResult(&ResponseSoftwareImageManagementSwimTaggingGoldenImage{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.TaggingGoldenImage(id, siteID, requestSoftwareImageManagementSwimTaggingGoldenImage)
		}

		return nil, response, fmt.Errorf("error with operation TaggingGoldenImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimTaggingGoldenImage)
	return result, response, err

}

//UntaggingGoldenImage Untagging golden image - 82a0-7b30-422a-b0ff
/* Untag the golden images specifically designed for a particular device type or supervisor engine module. Conditions for untagging the golden image:
1) Untagging the golden image can only be done where the golden tagged is applied.
  For example, if golden tagging is applied to a global site, then untagging can only be done on a global site. Even though the same setting will be inherited on native, attempting to untag will fail.
2) Untagging of SUBPACKAGE and ROMMON image type is not supported.


@param id id path parameter. Software image identifier is used for golden tagging or intent to tag it. The value of `id` can be obtained from the response of the API `/dna/intent/api/v1/images?imported=true&isAddonImages=false` for the base image and `/dna/images/{id}/addonImages` where `id` will be the software image identifier of the base image.

@param siteID siteId path parameter. Site identifier for tagged image or intent to tag it. The default value is global site id. See [https://developer.cisco.com/docs/dna-center](#!get-site) for `siteId`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!untagging-golden-image
*/
func (s *SoftwareImageManagementSwimService) UntaggingGoldenImage(id string, siteID string, requestSoftwareImageManagementSwimUntaggingGoldenImage *RequestSoftwareImageManagementSwimUntaggingGoldenImage) (*ResponseSoftwareImageManagementSwimUntaggingGoldenImage, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{id}/sites/{siteId}/untagGolden"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)
	path = strings.Replace(path, "{siteId}", fmt.Sprintf("%v", siteID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimUntaggingGoldenImage).
		SetResult(&ResponseSoftwareImageManagementSwimUntaggingGoldenImage{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.UntaggingGoldenImage(id, siteID, requestSoftwareImageManagementSwimUntaggingGoldenImage)
		}

		return nil, response, fmt.Errorf("error with operation UntaggingGoldenImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimUntaggingGoldenImage)
	return result, response, err

}

//AssignNetworkDeviceProductNameToTheGivenSoftwareImage Assign network device product name to the given software image - 0089-283d-4609-98a5
/* Assign network device product name and sites for the given image identifier. Refer `/dna/intent/api/v1/images` API for obtaining imageId


@param imageID imageId path parameter. Software image identifier. Refer `/dna/intent/api/v1/images` API for obtaining `imageId`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!assign-network-device-product-name-to-the-given-software-image
*/
func (s *SoftwareImageManagementSwimService) AssignNetworkDeviceProductNameToTheGivenSoftwareImage(imageID string, requestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImage *RequestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImage) (*ResponseSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImage, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{imageId}/siteWiseProductNames"
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImage).
		SetResult(&ResponseSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImage{}).
		SetError(&Error).
		Post(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {

		if response.StatusCode() == http.StatusUnauthorized {
			return s.AssignNetworkDeviceProductNameToTheGivenSoftwareImage(imageID, requestSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImage)
		}

		return nil, response, fmt.Errorf("error with operation AssignNetworkDeviceProductNameToTheGivenSoftwareImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimAssignNetworkDeviceProductNameToTheGivenSoftwareImage)
	return result, response, err

}

//UpdateRemoteImageDistributionServer Update remote image distribution server - 2caa-c9cc-469b-a3d5
/* Update remote image distribution server details.


@param id id path parameter. Remote server identifier.

*/
func (s *SoftwareImageManagementSwimService) UpdateRemoteImageDistributionServer(id string, requestSoftwareImageManagementSwimUpdateRemoteImageDistributionServer *RequestSoftwareImageManagementSwimUpdateRemoteImageDistributionServer) (*ResponseSoftwareImageManagementSwimUpdateRemoteImageDistributionServer, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/distributionServerSettings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimUpdateRemoteImageDistributionServer).
		SetResult(&ResponseSoftwareImageManagementSwimUpdateRemoteImageDistributionServer{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateRemoteImageDistributionServer(id, requestSoftwareImageManagementSwimUpdateRemoteImageDistributionServer)
		}
		return nil, response, fmt.Errorf("error with operation UpdateRemoteImageDistributionServer")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimUpdateRemoteImageDistributionServer)
	return result, response, err

}

//UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage Update the list of sites for the network device product name assigned to the software image - 1da6-d80b-40fa-8bdc
/* Update the list of sites for the network device product name assigned to the software image. Refer to `/dna/intent/api/v1/images` and `/dna/intent/api/v1/images/{imageId}/siteWiseProductNames` GET APIs for obtaining  `imageId` and `productNameOrdinal` respectively.


@param imageID imageId path parameter. Software image identifier. Refer `/dna/intent/api/v1/images` API for obtaining `imageId`

@param productNameOrdinal productNameOrdinal path parameter. Product name ordinal is unique value for each network device product. Refer `/dna/intent/api/v1/images/{imageId}/siteWiseProductNames` GET API for obtaining `productNameOrdinal`.

*/
func (s *SoftwareImageManagementSwimService) UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage(imageID string, productNameOrdinal float64, requestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage *RequestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage) (*ResponseSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage, *resty.Response, error) {
	path := "/dna/intent/api/v1/images/{imageId}/siteWiseProductNames/{productNameOrdinal}"
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)
	path = strings.Replace(path, "{productNameOrdinal}", fmt.Sprintf("%v", productNameOrdinal), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetBody(requestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage).
		SetResult(&ResponseSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage{}).
		SetError(&Error).
		Put(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage(imageID, productNameOrdinal, requestSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage)
		}
		return nil, response, fmt.Errorf("error with operation UpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimUpdateTheListOfSitesForTheNetworkDeviceProductNameAssignedToTheSoftwareImage)
	return result, response, err

}

//RemoveGoldenTagForImage Remove Golden Tag for image. - f3b9-5978-4f6a-897a
/* Remove golden tag. Set siteId as -1 for Global site.


@param siteID siteId path parameter. Site Id in uuid format. Set siteId as -1 for Global site.

@param deviceFamilyIDentifier deviceFamilyIdentifier path parameter. Device family identifier e.g. : 277696480-283933147, e.g. : 277696480

@param deviceRole deviceRole path parameter. Device Role. Permissible Values : ALL, UNKNOWN, ACCESS, BORDER ROUTER, DISTRIBUTION and CORE.

@param imageID imageId path parameter. Image Id in uuid format.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-golden-tag-for-image
*/
func (s *SoftwareImageManagementSwimService) RemoveGoldenTagForImage(siteID string, deviceFamilyIDentifier string, deviceRole string, imageID string) (*ResponseSoftwareImageManagementSwimRemoveGoldenTagForImage, *resty.Response, error) {
	//siteID string,deviceFamilyIDentifier string,deviceRole string,imageID string
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
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveGoldenTagForImage(siteID, deviceFamilyIDentifier, deviceRole, imageID)
		}
		return nil, response, fmt.Errorf("error with operation RemoveGoldenTagForImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRemoveGoldenTagForImage)
	return result, response, err

}

//RemoveImageDistributionServer Remove image distribution server - 43a7-5b17-404a-aa9f
/* Delete remote image distribution server.


@param id id path parameter. Remote server identifier.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!remove-image-distribution-server
*/
func (s *SoftwareImageManagementSwimService) RemoveImageDistributionServer(id string) (*ResponseSoftwareImageManagementSwimRemoveImageDistributionServer, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/images/distributionServerSettings/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimRemoveImageDistributionServer{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.RemoveImageDistributionServer(id)
		}
		return nil, response, fmt.Errorf("error with operation RemoveImageDistributionServer")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimRemoveImageDistributionServer)
	return result, response, err

}

//DeleteImage Delete image - 0c82-1b84-44da-8bc0
/* Delete the image from image repository


@param id id path parameter. The software image identifier that needs to be deleted can be obtained from the API `/dna/intent/api/v1/images?imported=true`. Use this API to obtain the `id` of the image.


Documentation Link: https://developer.cisco.com/docs/dna-center/#!delete-image
*/
func (s *SoftwareImageManagementSwimService) DeleteImage(id string) (*ResponseSoftwareImageManagementSwimDeleteImage, *resty.Response, error) {
	//id string
	path := "/dna/intent/api/v1/images/{id}"
	path = strings.Replace(path, "{id}", fmt.Sprintf("%v", id), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimDeleteImage{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.DeleteImage(id)
		}
		return nil, response, fmt.Errorf("error with operation DeleteImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimDeleteImage)
	return result, response, err

}

//UnassignNetworkDeviceProductNameFromTheGivenSoftwareImage Unassign network device product name from the given software image - 3fa4-39e3-4a4b-8eaf
/* This API unassigns the network device product name from all the sites for the given software image.
        Refer to `/dna/intent/api/v1/images` and `/dna/intent/api/v1/images/{imageId}/siteWiseProductNames` GET APIs for obtaining  `imageId` and `productNameOrdinal` respectively.


@param imageID imageId path parameter. Software image identifier. Refer `/dna/intent/api/v1/images` API for obtaining `imageId`

@param productNameOrdinal productNameOrdinal path parameter. The product name ordinal is a unique value for each network device product. Refer `/dna/intent/api/v1/images/{imageId}/siteWiseProductNames` GET API for obtaining `productNameOrdinal`


Documentation Link: https://developer.cisco.com/docs/dna-center/#!unassign-network-device-product-name-from-the-given-software-image
*/
func (s *SoftwareImageManagementSwimService) UnassignNetworkDeviceProductNameFromTheGivenSoftwareImage(imageID string, productNameOrdinal float64) (*ResponseSoftwareImageManagementSwimUnassignNetworkDeviceProductNameFromTheGivenSoftwareImage, *resty.Response, error) {
	//imageID string,productNameOrdinal float64
	path := "/dna/intent/api/v1/images/{imageId}/siteWiseProductNames/{productNameOrdinal}"
	path = strings.Replace(path, "{imageId}", fmt.Sprintf("%v", imageID), -1)
	path = strings.Replace(path, "{productNameOrdinal}", fmt.Sprintf("%v", productNameOrdinal), -1)

	response, err := s.client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetResult(&ResponseSoftwareImageManagementSwimUnassignNetworkDeviceProductNameFromTheGivenSoftwareImage{}).
		SetError(&Error).
		Delete(path)

	if err != nil {
		return nil, nil, err

	}

	if response.IsError() {
		if response.StatusCode() == http.StatusUnauthorized {
			return s.UnassignNetworkDeviceProductNameFromTheGivenSoftwareImage(imageID, productNameOrdinal)
		}
		return nil, response, fmt.Errorf("error with operation UnassignNetworkDeviceProductNameFromTheGivenSoftwareImage")
	}

	result := response.Result().(*ResponseSoftwareImageManagementSwimUnassignNetworkDeviceProductNameFromTheGivenSoftwareImage)
	return result, response, err

}
