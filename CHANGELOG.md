# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

## [2.0.0] - 2021-11-30
### Added
- Add function to get the resty.Client pointer.

### Changed
- Update to use DNAC 2.2.2.3

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
[2.0.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v1.0.0...v2.0.0
[2.1.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v2.0.0...v2.1.0
[3.0.0]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v2.0.0...v3.0.0
[Unreleased]: https://github.com/cisco-en-programmability/dnacenter-go-sdk/compare/v3.0.0...main
