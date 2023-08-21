# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [5.0.9] - 2023-08-21

### Changed
-  `ResponseSitesGetSiteResponseAdditionalInfoAttributes` is transformed into a structure with the following attributes:
   - Country `string`
   - Address `string`
   - Latitude `string`
   - AddressInheritedFrom `string`
   - Type `string`
   - Longitude `string`
   - OffsetX `string`
   - OffsetY `string`
   - Length `string`
   - Width `string`
   - Height `string`
   - RfModel `string`
   - FloorIndex `string`

## [5.0.8] - 2023-07-31

### Changed
- `ResponseSdaGetDefaultAuthenticationProfileFromSdaFabric.WakeOnLan` change from `string` to `interface`. To consider both `string` and `bool` response values.
- `GetDefaultAuthenticationProfileFromSdaFabric` changes empty string response to `nil`

## [5.0.7] - 2023-07-12
### Changed
- `RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfo` struct is now a pointer.
- `ConfigParameters` struct is now a pointer.
- `Skip` is a pointer.

## [5.0.6] - 2023-06-26
### Changed
- Device onboarding Pnp Service Changes `ConfigInfo` change to `object` on `RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfo` struct.
- SDA Service Changes `WakeOnLan` change to `string` on `ResponseSdaGetDefaultAuthenticationProfileFromSdaFabric` struct.

## [5.0.5] - 2023-05-11
### Changed
- Device onboarding Pnp Service Changes `ConfigParameters` change to `array` on `RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfo` struct.

## [5.0.4] - 2023-05-05
### Changed
- SDA Service Changes `IsInfraVn` and `IsDefaultVn` change to `*bool` on `ResponseSdaGetVnFromSdaFabric` struct.

## [5.0.3] - 2023-05-03
### Changed
- ResponseSdaGetDefaultAuthenticationProfileFromSdaFabric is an object not an array.
- ResponseNetworkSettingsGetGlobalPoolResponse adding new attributes and changing attributes types.

## [5.0.2] - 2023-04-27
### Changed
- Adding struct for sites.ResponseSitesGetSiteResponse.AdditionalInfo

## [5.0.1] - 2023-04-12
### Changed
- Mulicast -> Multicast on RequestItemLanAutomationLanAutomationStart

## [5.0.0] - 2023-04-12
### Added
- New services for Cisco DNA Center 2.3.5.3's API:
  - AuthenticationManagement *AuthenticationManagementService
  - DisasterRecovery *DisasterRecoveryService
  - EoX *EoXService
  - ItsmIntegration *ItsmIntegrationService
  - *PlatformConfigurationService turns to *PlatformService
  - Policy *PolicyService
  - UserandRoles *UserandRolesService
- New API Methods for previous Services:
  - DevicesService.CreateUserDefinedField
  - DevicesService.UpdateUserDefinedField
  - DevicesService.AddUserDefinedFieldToDevice
  - DevicesService.DeleteUserDefinedField
  - DevicesService.RemoveUserDefinedFieldFromDevice
  - DiscoveryService.GetAllGlobalCredentialsV2
  - DiscoveryService.CreateGlobalCredentialsV2
  - DiscoveryService.UpdateGlobalCredentialsV2
  - DiscoveryService.DeleteGlobalCredentialV2
  - EventManagementService.GetEmailDestination
  - EventManagementService.GetSNMPDestination
  - EventManagementService.GetSyslogDestination
  - EventManagementService.GetWebhookDestination
  - FabricWirelessService.AddSSIDToIPPoolMapping
  - IssuesService.ExecuteSuggestedActionsCommands
  - LanAutomationService.LanAutomationLogsForIndividualDevices
  - LanAutomationService.LanAutomationActiveSessions
  - NetworkSettingsService.GetNetworkV2
  - NetworkSettingsService.GetServiceProviderDetailsV2
  - NetworkSettingsService.AssignDeviceCredentialToSiteV2
  - NetworkSettingsService.CreateNetworkV2
  - NetworkSettingsService.CreateSpProfileV2
  - NetworkSettingsService.UpdateNetworkV2
  - NetworkSettingsService.UpdateSpProfileV2
  - NetworkSettingsService.DeleteSpProfileV2
  - SystemSettingsService.GetAuthenticationAndPolicyServers
  - WirelessService.GetAccessPointRebootTaskResult
  - WirelessService.GetAccessPointConfigurationTaskResult
  - WirelessService.GetAccessPointConfiguration
  - WirelessService.RebootAccessPoints
  - WirelessService.ConfigureAccessPoints

- The SDK was updated with the official DNA Center API 2.3.5.3 documentation, therefore some structures or data types may have changed.[Oficial Documentation](https://developer.cisco.com/docs/dna-center/#!cisco-dna-center-2-3-5-api-overview).



## [4.0.14] - 2023-02-27
- `ResponseClientsGetClientDetailDetail.Detail.VlanId` atributte change type from `string` to `*int` in `clients` service.

## [4.0.13] - 2023-02-02
- ID attribute added to `ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessList` on `device_onboarding_pnp` service.

## [4.0.12] - 2023-01-10
- Query params `offset` and `limit` change to `int`.

## [4.0.11] - 2022-11-16
- Change `ConnectedToInternet` and `BorderWithExternalConnectivity` parameters from `bool` to `string`.
  
## [4.0.10] - 2022-11-15
- Change`RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL2Handoff` to allow array.
  
## [4.0.9] - 2022-11-14
### Changed
- Change `RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3Handoff` to allow array.

## [4.0.8] - 2022-10-21
### Added
- Adding `ResponseSiteGetSite` response struct.

## [4.0.7] - 2022-10-18
### Changed
- Change `ResponseSitesDeleteSite` response struct.

## [4.0.6] - 2022-07-19
### Added
- Adding `GetSiteByID` method.

## [4.0.5] - 2022-07-19
### Changed
- Adding `to many request` handle. New parameter on configuration options of SDK API Client.
- Method `SetDNACWaitTimeToManyRequest` added for update wait time (in minutes) to reintent API requests that fails due `429` status code.
- Adding `GetSiteByID` method.

## [4.0.4] - 2022-07-19
### Changed
- `sda.ResponseSdaGetBorderDeviceDetailFromSdaFabric.Payload` atributte removed from `sda.ResponseSdaGetBorderDeviceDetailFromSdaFabric`.
  
## [4.0.3] - 2022-07-12
### Changed
- `[]RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3HandoffVirtualNetwork` becomes to `RequestItemSdaAddBorderDeviceInSdaFabricExternalConnectivitySettingsL3HandoffVirtualNetwork`.
  
## [4.0.2] - 2022-07-08
### Changed
- `event_management.ResponseItemEventManagementGetEventSubscriptionsSubscriptionEndpointsSubscriptionDetails.TrustCert` parameter turns to `boolean`.
- `event_management.ResponseItemEventManagementGetEventSubscriptionsSubscriptionEndpointsSubscriptionDetails.ConnectTimeout` parameter turns to `number`.
- `event_management.ResponseItemEventManagementGetEventSubscriptionsSubscriptionEndpointsSubscriptionDetails.ReadTimeout` parameter turns to `number`.
  
### Added
- Added `SiteNameHierarchy`, `FabricName`, `FabricType` and `FabricDomainType` to
`ResponseSdaGetSiteFromSdaFabric` of `sda`

## [4.0.1] - 2022-06-17
### Changed
- `discovery.ResponseDiscoveryGetGlobalCredentialsResponse.Secure` parameter turns to `boolean`.
- `event_management.ResponseItemEventManagementGetSyslogEventSubscriptionsSubscriptionEndpointsSubscriptionDetailsSyslogConfig.Port` turns to `int`.

## [4.0.0] - 2022-06-16

### Added
New services for Cisco DNA Center 2.3.3.0's API:
- CiscoDnaCenterSystem *CiscoDnaCenterSystemService
- LanAutomation *LanAutomationService
- SystemSettings *SystemSettingsService

### Removed
Services removed on Cisco DNA Center 2.3.3.0's API:
- AuthenticationManagement *AuthenticationManagementService
- DisasterRecovery *DisasterRecoveryService
- Policy *PolicyService
  
### Changed
- SDK now compatible with Cisco DNA Center 2.3.3.0's API.
  
## [3.6.3] - 2022-03-17
### Changed
- Changed `RequestConfigurationTemplatesDeployTemplateV2TargetInfoParams` from `interface` to `map[string]interface{}`
## [3.6.2] - 2022-03-17
### Changed
- Changed `ResponseDevicesGetDeviceBySerialNumberResponse.LastUpdateTime` from `string` to `*int`
- Changed `RequestConfigurationTemplatesDeployTemplateTargetInfoParams` from `interface` to `map[string]interface{}`
## [3.6.1] - 2022-03-16
### Removed
- Removed from `RequestDeviceOnboardingPnpClaimADeviceToASite` following pointers.
    + `RequestDeviceOnboardingPnpClaimADeviceToASiteImageInfo`.
    + `RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfo`.
- Removed from `RequestDeviceOnboardingPnpClaimADeviceToASiteImageInfo` following pointers.
    + `Skip`.
- Removed from `RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfo` following pointers.
    + `RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfoConfigParameters`.
- Removed `omitempty` for following variables.
    + `RequestDeviceOnboardingPnpClaimADeviceToASite.ImageInfo`
    + `RequestDeviceOnboardingPnpClaimADeviceToASite.ConfigInfo`
    + `RequestDeviceOnboardingPnpClaimADeviceToASite.ImageInfo.ImageID`
    + `RequestDeviceOnboardingPnpClaimADeviceToASite.ImageInfo.Skip`
    + `RequestDeviceOnboardingPnpClaimADeviceToASite.ConfigInfo.ConfigID`
    + `RequestDeviceOnboardingPnpClaimADeviceToASite.ConfigInfo.ConfigParameters`
    + `RequestDeviceOnboardingPnpClaimADeviceToASite.ConfigInfo.ConfigParameters.Key`
    + `RequestDeviceOnboardingPnpClaimADeviceToASite.ConfigInfo.ConfigParameters.Value`
  
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
[3.6.1]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.6.0...v3.6.1
[3.6.2]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.6.1...v3.6.2
[3.6.3]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.6.2...v3.6.3
[4.0.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.6.3...v4.0.0
[4.0.1]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.0...v4.0.1
[4.0.2]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.1...v4.0.2
[4.0.3]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.2...v4.0.3
[4.0.4]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.3...v4.0.4
[4.0.5]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.4...v4.0.5
[4.0.6]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.5...v4.0.6
[4.0.7]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.6...v4.0.7
[4.0.8]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.7...v4.0.8
[4.0.9]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.8...v4.0.9
[4.0.10]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.9...v4.0.10
[4.0.11]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.10...v4.0.11
[4.0.12]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.11...v4.0.12
[4.0.13]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.12...v4.0.13
[4.0.14]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.13...v4.0.14
[5.0.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v4.0.14...v5.0.0
[5.0.1]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v5.0.0...v5.0.1
[5.0.2]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v5.0.1...v5.0.2
[5.0.3]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v5.0.2...v5.0.3
[5.0.4]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v5.0.3...v5.0.4
[5.0.5]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v5.0.4...v5.0.5
[5.0.6]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v5.0.5...v5.0.6
[5.0.7]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v5.0.6...v5.0.7
[5.0.8]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v5.0.7...v5.0.8
[Unreleased]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v5.0.8...main
