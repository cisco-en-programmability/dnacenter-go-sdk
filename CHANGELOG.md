# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
## [3.6.0] - 2022-03-15
### Added
- Added new service `CustomCallService`
## [3.5.1] - 2022-03-1
### Added
- Added new struct `ResponseWirelessPSKOverride`  on `wireless`for `PSKOverride` response.
### Changed
- Changed type of following variable on `RequestTagUpdatesTagMembership` struct of `tag`.
    + `MemberToTags` from []RequestTagUpdatesTagMembershipMemberToTags to map[string][]string.
### Removed
- Removed `RequestTagUpdatesTagMembershipMemberToTags` struct from `tag`.

## [3.5.0] - 2022-02-17
### Added
- Added `Password`, `Port`, `Secure`, `Username`, `NetconfPort`, `ReadCommunity` and `WriteCommunity` to
`ResponseDiscoveryGetGlobalCredentialsResponse` of `discovery`
### Changed
- Changed type of the following variable on `ResponseNetworkSettingsGetServiceProviderDetailsResponse` struct of `network-settings`
    + `Version` from string to int
- Changed type of the following variable on `ResponseSitesGetSiteResponse` struct of `sites`
    + `AdditionalInfo` from string to `[]ResponseSitesGetSiteResponseAdditionalInfo`

## [3.4.1] - 2022-02-01
### Added
- Added ImageInfo, ConfigInfo and Hostname to RequestDeviceOnboardingPnpClaimADeviceToASite struct in device_onboarding_pnp file.

### Changed
- ResponseHealthAndPerformanceSystemPerformanceHistoricalAPIKpisData changes from struct with `T1` property to a `map[string][]string`

## [3.4.0] - 2022-01-27
### Added
- Added SubscriptionID, IsPrivate, TenantID to ResponseItemEventManagementGetEventSubscriptions of event_management.
- Added IsDeletable to ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProject of configutations_templates.
- Add ID to RequestDevicesSyncDevices2

### Removed
- Removes ResponseItemFabricWirelessUpdateSSIDToIPPoolMapping, ResponseItemFabricWirelessUpdateSSIDToIPPoolMapping,ResponseItemFabricWirelessAddWLCToFabricDomain from fabric_wireless.
- Removes ResponseItemWirelessApProvision, ResponseItemWirelessCreateUpdateDynamicInterface from wireless.

### Changed

- Changes ResponseFabricWirelessAddSSIDToIPPoolMapping to prev ResponseItemFabricWirelessAddSSIDToIPPoolMapping, ResponseFabricWirelessUpdateSSIDToIPPoolMapping to prev ResponseItemFabricWirelessUpdateSSIDToIPPoolMapping, ResponseFabricWirelessAddWLCToFabricDomain to prev ResponseItemFabricWirelessAddWLCToFabricDomain, ResponseDeviceOnboardingPnpGetDeviceList2 to prev ResponseItemDeviceOnboardingPnpGetDeviceList2.
- Changes ResponseWirelessApProvision to prev ResponseItemWirelessApProvision.
- Changes ResponseWirelessCreateUpdateDynamicInterface to prev ResponseItemWirelessCreateUpdateDynamicInterface.

## [3.3.1] - 2022-01-19

### Added
- Creation of a new `ResponseApplicationPolicyGetApplicationsResponse` structure.

### Changed
- On `Network Settings` changed the type of the following variables:
    + `CreateTime`            from string to int
    + `LastUpdateTime`        from string to int
    + `TotalIPAddressCount`   from string to int 64
    + `UsedIPAddressCount`    from string to int 64
- Changed type of the following variables on `ResponseApplicationPolicyGetApplicationsResponseNetworkApplications` struct
    + `Popularity`  from string to int
    + `Rank`        from string to int
- Changed type of the following variables on `ResponseApplicationPolicyGetApplicationsResponseNetworkIDentity` struct
    + `LowertPort`    from string to int
    + `UpperPort`     from string to int

## [3.3.0] - 2022-01-13

### Added
- application_policy - Adds IndicativeNetworkIdentity property to ResponseItemApplicationPolicyGetApplications struct
- application_policy - Adds ResponseItemApplicationPolicyGetApplicationsIndicativeNetworkIDentity struct

### Fixed
- Fix SetAuthToken to use DNAC X-Auth-Token
- sites - Response of ResponseSitesGetSiteCount type changes from string to *int

## [3.2.0] - 2022-01-13

### Added
- Adds "Separating Authentication" [#17](https://github.com/cisco-en-programmability/dnacenter-go-sdk/pull/17):
    + Adds `SetOptions` function.
    + Adds `NewClientNoAuth` function.
    + Adds `NewClientWithOptionsNoAuth` function.
    + Adds `AuthClient ` function.

### Fixed
- Fixes "Add ability to test via Resty httpmock" [#10](https://github.com/cisco-en-programmability/dnacenter-go-sdk/issues/10)
- Changes type from int to float64 for ResponseTopologyGetOverallNetworkHealthHealthDistirubution.GoodPercentage in topology.go

## [3.1.1] - 2022-01-10

### Added
- Adds `FabricName` property to `ResponseSdaGetSdaFabricInfo`
- Adds `FabricType` property to `ResponseSdaGetSdaFabricInfo`
- Adds `FabricDomainType` property to `ResponseSdaGetSdaFabricInfo`

### Changed
- `ExecutionStatusURL` changes to `ExecutionId` on `ResponseSdaGetSdaFabricInfo`
## [3.1.0] - 2021-12-24

### Added
- Adds `RequestSiteDesignCreateFloormap` struct
- Adds `RequestSiteDesignUpdateFloormap` struct
- Adds `requestSiteDesignCreateFloormap` param to `CreateFloormap`
- Adds `requestSiteDesignUpdateFloormap` param to `UpdateFloormap`

## [3.0.2] - 2021-12-14
### Fixed
- Fixes ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttribute struct
- Fixes RequestTagAddMembersToTheTag struct

## [3.0.1] - 2021-12-13
### Fixed
- Fix ImportCertificateP12

## [3.0.0] - 2021-12-06
### Added
- Add ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplates struct (change interface{})
- Add ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplates struct (change interface{})
- Add DocumentDatabase and ProjectAssociated properties to ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplate
- Add ImportCertificateMultipartFields to ImportCertificate
- Add ImportCertificateP12MultipartFields to ImportCertificateP12
- Add VirtualNetworkType to RequestSdaAddVirtualNetworkWithScalableGroups and RequestSdaUpdateVirtualNetworkWithScalableGroups
- Add ImportLocalSoftwareImageMultipartFields to ImportLocalSoftwareImage
- Add BapiError to ResponseTaskGetBusinessAPIExecutionDetails

### Removed
- Removes CreatesACloneOfTheGivenTemplateQueryParams

### Changed
- Update to use DNAC 2.2.3.3
- Changes type of Templates property from (pointer to struct) to (pointer to array of struct):
    + ResponseItemConfigurationTemplatesGetsAListOfProjects
    + ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProject
    + ResponseConfigurationTemplatesGetProjectsDetailsResponse
    + RequestConfigurationTemplatesCreateProject
    + RequestConfigurationTemplatesUpdateProject
- Changes type of ResourceParams property from (pointer to struct) to (pointer to array of struct):
    + RequestConfigurationTemplatesDeployTemplateTargetInfo
    + RequestConfigurationTemplatesPreviewTemplate
    + RequestConfigurationTemplatesDeployTemplateV2TargetInfo
- Change type of RollbackTemplateErrors TemplateErrors and properties from (pointer to struct) to (pointer to array of struct):
    + ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrors
    + ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplatesValidationErrors
    + ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateValidationErrors
    + ResponseConfigurationTemplatesGetProjectsDetailsResponseTemplatesValidationErrors
    + ResponseConfigurationTemplatesGetTemplatesDetailsResponseValidationErrors
    + RequestConfigurationTemplatesCreateProjectTemplatesValidationErrors
    + RequestConfigurationTemplatesUpdateProjectTemplatesValidationErrors
    + RequestItemConfigurationTemplatesImportsTheTemplatesProvidedValidationErrors
    + RequestConfigurationTemplatesCreateTemplateValidationErrors
    + RequestConfigurationTemplatesUpdateTemplateValidationErrors
- Change GetProjectsDetails, GetTemplatesDetails response struct to ResponseConfigurationTemplatesGetProjectsDetails, ResponseConfigurationTemplatesGetTemplatesDetails (which adds response property with a list of previous struct definition)

## [2.2.1] - 2022-01-13

### Fixed
- Fix SetAuthToken to use DNAC X-Auth-Token

## [2.2.0] - 2022-01-13

### Added
- Adds "Separating Authentication" [#17](https://github.com/cisco-en-programmability/dnacenter-go-sdk/pull/17):
    + Adds `SetOptions` function.
    + Adds `NewClientNoAuth` function.
    + Adds `NewClientWithOptionsNoAuth` function.
    + Adds `AuthClient ` function.

## [2.1.2] - 2021-12-14
### Fixed
- Fixes ResponseDevicesGetDeviceValuesThatMatchFullyOrPartiallyAnAttribute struct
- Fixes RequestTagAddMembersToTheTag struct

## [2.1.1] - 2021-12-13
### Fixed
- Fix ImportCertificateP12

## [2.1.0] - 2021-12-06
### Added
- Add ImportCertificateMultipartFields to ImportCertificate
- Add ImportCertificateP12MultipartFields to ImportCertificateP12
- Add ImportLocalSoftwareImageMultipartFields to ImportLocalSoftwareImage
## [2.0.0] - 2021-12-06
### Added
- Add function to get the resty.Client pointer.

### Changed
- Update to use DNAC 2.2.2.3

## [1.2.1] - 2022-01-13

### Fixed
- Fix SetAuthToken to use DNAC X-Auth-Token

## [1.2.0] - 2022-01-13

### Added
- Adds "Separating Authentication" [#17](https://github.com/cisco-en-programmability/dnacenter-go-sdk/pull/17):
    + Adds `SetOptions` function.
    + Adds `NewClientNoAuth` function.
    + Adds `NewClientWithOptionsNoAuth` function.
    + Adds `AuthClient ` function.
## [1.1.0] - 2021-12-06

### Added
- Add ImportCertificateMultipartFields to ImportCertificate
- Add ImportCertificateP12MultipartFields to ImportCertificateP12
- Add ImportLocalSoftwareImageMultipartFields to ImportLocalSoftwareImage
## [1.0.0] - 2021-11-30
### Added
- Common client added for Resty.Client reference.

### Changed
- Update issue templates.

### Removed
- Removed `var RestyClient *resty.Client`

### Fixed
- Solved issue 'Add ability to test via Resty httpmock'

## [0.1.1] - 2021-11-02

### Changed
- Int for count values and float64 for percentages in topology.go
  
### Fixed
- Int for count values and float64 for percentages in topology.go

## [0.1.0] - 2021-10-29
### Added
- Adds error response to api_client.go
- Adds Delete__Response to Delete calls for sites and tag
- Adds check response.IsError() to get calls in sites and tag
- New definition for GetApplicationsResponse, uses GetApplicationsResponseResponse
- Adds Attributes to GetSiteResponseResponseAdditionalInfo
- Adds GetSiteResponseResponseAdditionalInfoAttributes struct
- Add headers for GetClientEnrichmentDetails 
  
### Changed
- Changes structure definitions of application_policy and sites
- Changes struct name from GetApplicationsResponse to GetApplicationsResponseResponse
- AdditionalInfo is []GetSiteResponseResponseAdditionalInfo, not GetSiteResponseResponseAdditionalInfo
- Changes json attribute, replaces namespace to nameSpace
- Changes the type in network_settings.go from int to uint64 for GetGlobalPoolResponseResponse
- Update README 
- Update example reference
- Update dependencies
- Add workflow and scripts

### Fixed
- Fixes GetApplicationsResponse struct
- Fixes GetSiteResponseResponseAdditionalInfo struct
- AdditionalInfo is []GetSiteResponseResponseAdditionalInfo, not GetSiteResponseResponseAdditionalInfo
- Fixes Task structs
- Fixes GetSiteCountResponse struct
- Fixes ConfigurationTemplates struct
- Fixes Discovery struct
- Fixes device_onboarding_pnp structures and return values
- Fixes network_settings structures
- Fixes some sites structures
- Fixes sda's GetSDAFabricInfoResponse struct
- Fix unchecked errors
- Fix test errors

[0.1.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/commits/v0.1.0
[0.1.1]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v0.1.0...v0.1.1
[1.0.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v0.1.1...v1.0.0
[1.1.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v1.0.0...v1.1.0
[1.2.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v1.1.0...v1.2.0
[1.2.1]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v1.2.0...v1.2.1
[2.0.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v1.0.0...v2.0.0
[2.1.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v2.0.0...v2.1.0
[2.1.1]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v2.1.0...v2.1.1
[2.1.2]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v2.1.1...v2.1.2
[2.2.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v2.1.2...v2.2.0
[2.2.1]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v2.2.0...v2.2.1
[3.0.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v2.0.0...v3.0.0
[3.0.1]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.0.0...v3.0.1
[3.0.2]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.0.1...v3.0.2
[3.1.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.0.2...v3.1.0
[3.1.1]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.1.0...v3.1.1
[3.2.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.1.1...v3.2.0
[3.3.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.2.0...v3.3.0
[3.3.1]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.3.0...v3.3.1
[3.4.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.3.1...v3.4.0
[3.4.1]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.4.0...v3.4.1
[3.5.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.4.1...v3.5.0
[3.5.1]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.5.0...v3.5.1
[3.6.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.5.1...v3.6.0
[Unreleased]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.6.0...main
