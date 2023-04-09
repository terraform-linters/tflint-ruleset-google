// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package magicmodules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// Rules is a list of rules generated from Magic Modules
var Rules = []tflint.Rule{
	NewGoogleAccessContextManagerAuthorizedOrgsDescInvalidAssetTypeRule(),
	NewGoogleAccessContextManagerAuthorizedOrgsDescInvalidAuthorizationDirectionRule(),
	NewGoogleAccessContextManagerAuthorizedOrgsDescInvalidAuthorizationTypeRule(),
	NewGoogleAccessContextManagerServicePerimeterInvalidPerimeterTypeRule(),
	NewGoogleActiveDirectoryDomainTrustInvalidTrustDirectionRule(),
	NewGoogleActiveDirectoryDomainTrustInvalidTrustTypeRule(),
	NewGoogleAlloydbInstanceInvalidAvailabilityTypeRule(),
	NewGoogleAlloydbInstanceInvalidInstanceTypeRule(),
	NewGoogleApigeeEnvironmentInvalidApiProxyTypeRule(),
	NewGoogleApigeeEnvironmentInvalidDeploymentTypeRule(),
	NewGoogleApigeeOrganizationInvalidRetentionRule(),
	NewGoogleApigeeOrganizationInvalidRuntimeTypeRule(),
	NewGoogleAppEngineDomainMappingInvalidOverrideStrategyRule(),
	NewGoogleAppEngineFirewallRuleInvalidActionRule(),
	NewGoogleAppEngineFlexibleAppVersionInvalidServingStatusRule(),
	NewGoogleArtifactRegistryRepositoryInvalidModeRule(),
	NewGoogleBeyondcorpAppGatewayInvalidHostTypeRule(),
	NewGoogleBeyondcorpAppGatewayInvalidTypeRule(),
	NewGoogleBigQueryRoutineInvalidDeterminismLevelRule(),
	NewGoogleBigQueryRoutineInvalidLanguageRule(),
	NewGoogleBigQueryRoutineInvalidRoutineTypeRule(),
	NewGoogleBigqueryDatapolicyDataPolicyInvalidDataPolicyTypeRule(),
	NewGoogleBinaryAuthorizationPolicyInvalidGlobalPolicyEvaluationModeRule(),
	NewGoogleCloudAssetFolderFeedInvalidContentTypeRule(),
	NewGoogleCloudAssetOrganizationFeedInvalidContentTypeRule(),
	NewGoogleCloudAssetProjectFeedInvalidContentTypeRule(),
	NewGoogleCloudBuildTriggerInvalidIncludeBuildLogsRule(),
	NewGoogleCloudIdentityGroupInvalidInitialGroupConfigRule(),
	NewGoogleCloudIdsEndpointInvalidSeverityRule(),
	NewGoogleCloudIotDeviceInvalidLogLevelRule(),
	NewGoogleCloudRunV2JobInvalidLaunchStageRule(),
	NewGoogleCloudRunV2ServiceInvalidIngressRule(),
	NewGoogleCloudRunV2ServiceInvalidLaunchStageRule(),
	NewGoogleCloudiotRegistryInvalidLogLevelRule(),
	NewGoogleComputeAddressInvalidAddressTypeRule(),
	NewGoogleComputeAddressInvalidNameRule(),
	NewGoogleComputeAddressInvalidNetworkTierRule(),
	NewGoogleComputeBackendBucketInvalidCompressionModeRule(),
	NewGoogleComputeBackendBucketInvalidNameRule(),
	NewGoogleComputeBackendBucketSignedUrlKeyInvalidNameRule(),
	NewGoogleComputeBackendServiceInvalidCompressionModeRule(),
	NewGoogleComputeBackendServiceInvalidLoadBalancingSchemeRule(),
	NewGoogleComputeBackendServiceInvalidLocalityLbPolicyRule(),
	NewGoogleComputeBackendServiceInvalidProtocolRule(),
	NewGoogleComputeBackendServiceInvalidSessionAffinityRule(),
	NewGoogleComputeBackendServiceSignedUrlKeyInvalidNameRule(),
	NewGoogleComputeExternalVpnGatewayInvalidRedundancyTypeRule(),
	NewGoogleComputeFirewallInvalidDirectionRule(),
	NewGoogleComputeForwardingRuleInvalidIpProtocolRule(),
	NewGoogleComputeForwardingRuleInvalidLoadBalancingSchemeRule(),
	NewGoogleComputeForwardingRuleInvalidNetworkTierRule(),
	NewGoogleComputeGlobalAddressInvalidAddressTypeRule(),
	NewGoogleComputeGlobalAddressInvalidIpVersionRule(),
	NewGoogleComputeGlobalForwardingRuleInvalidIpProtocolRule(),
	NewGoogleComputeGlobalForwardingRuleInvalidIpVersionRule(),
	NewGoogleComputeGlobalForwardingRuleInvalidLoadBalancingSchemeRule(),
	NewGoogleComputeGlobalNetworkEndpointGroupInvalidNetworkEndpointTypeRule(),
	NewGoogleComputeHaVpnGatewayInvalidStackTypeRule(),
	NewGoogleComputeInterconnectAttachmentInvalidBandwidthRule(),
	NewGoogleComputeInterconnectAttachmentInvalidEncryptionRule(),
	NewGoogleComputeInterconnectAttachmentInvalidNameRule(),
	NewGoogleComputeInterconnectAttachmentInvalidTypeRule(),
	NewGoogleComputeManagedSslCertificateInvalidTypeRule(),
	NewGoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule(),
	NewGoogleComputeNetworkInvalidNetworkFirewallPolicyEnforcementOrderRule(),
	NewGoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule(),
	NewGoogleComputeRegionBackendServiceInvalidLoadBalancingSchemeRule(),
	NewGoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule(),
	NewGoogleComputeRegionBackendServiceInvalidProtocolRule(),
	NewGoogleComputeRegionBackendServiceInvalidSessionAffinityRule(),
	NewGoogleComputeRegionNetworkEndpointGroupInvalidNetworkEndpointTypeRule(),
	NewGoogleComputeRegionTargetTcpProxyInvalidProxyHeaderRule(),
	NewGoogleComputeRouteInvalidNameRule(),
	NewGoogleComputeRouterNatInvalidNatIpAllocateOptionRule(),
	NewGoogleComputeRouterNatInvalidSourceSubnetworkIpRangesToNatRule(),
	NewGoogleComputeRouterPeerInvalidAdvertiseModeRule(),
	NewGoogleComputeSslPolicyInvalidMinTlsVersionRule(),
	NewGoogleComputeSslPolicyInvalidProfileRule(),
	NewGoogleComputeSubnetworkInvalidIpv6AccessTypeRule(),
	NewGoogleComputeSubnetworkInvalidRoleRule(),
	NewGoogleComputeSubnetworkInvalidStackTypeRule(),
	NewGoogleComputeTargetHttpsProxyInvalidQuicOverrideRule(),
	NewGoogleComputeTargetInstanceInvalidNatPolicyRule(),
	NewGoogleComputeTargetSslProxyInvalidProxyHeaderRule(),
	NewGoogleComputeTargetTcpProxyInvalidProxyHeaderRule(),
	NewGoogleDataCatalogEntryGroupInvalidEntryGroupIdRule(),
	NewGoogleDataCatalogEntryInvalidTypeRule(),
	NewGoogleDataCatalogEntryInvalidUserSpecifiedSystemRule(),
	NewGoogleDataCatalogEntryInvalidUserSpecifiedTypeRule(),
	NewGoogleDataCatalogTagTemplateInvalidTagTemplateIdRule(),
	NewGoogleDataFusionInstanceInvalidTypeRule(),
	NewGoogleDataLossPreventionJobTriggerInvalidStatusRule(),
	NewGoogleDataprocMetastoreServiceInvalidDatabaseTypeRule(),
	NewGoogleDataprocMetastoreServiceInvalidReleaseChannelRule(),
	NewGoogleDataprocMetastoreServiceInvalidTierRule(),
	NewGoogleDatastoreIndexInvalidAncestorRule(),
	NewGoogleDeploymentManagerDeploymentInvalidCreatePolicyRule(),
	NewGoogleDeploymentManagerDeploymentInvalidDeletePolicyRule(),
	NewGoogleDialogflowAgentInvalidApiVersionRule(),
	NewGoogleDialogflowAgentInvalidMatchModeRule(),
	NewGoogleDialogflowAgentInvalidTierRule(),
	NewGoogleDialogflowCxEntityTypeInvalidAutoExpansionModeRule(),
	NewGoogleDialogflowCxEntityTypeInvalidKindRule(),
	NewGoogleDialogflowEntityTypeInvalidKindRule(),
	NewGoogleDialogflowIntentInvalidWebhookStateRule(),
	NewGoogleDnsManagedZoneInvalidVisibilityRule(),
	NewGoogleFirestoreDatabaseInvalidAppEngineIntegrationModeRule(),
	NewGoogleFirestoreDatabaseInvalidConcurrencyModeRule(),
	NewGoogleFirestoreDatabaseInvalidTypeRule(),
	NewGoogleFirestoreIndexInvalidQueryScopeRule(),
	NewGoogleHealthcareFhirStoreInvalidVersionRule(),
	NewGoogleKmsCryptoKeyInvalidPurposeRule(),
	NewGoogleKmsCryptoKeyVersionInvalidStateRule(),
	NewGoogleKmsKeyRingImportJobInvalidImportMethodRule(),
	NewGoogleKmsKeyRingImportJobInvalidProtectionLevelRule(),
	NewGoogleMemcacheInstanceInvalidMemcacheVersionRule(),
	NewGoogleMonitoringAlertPolicyInvalidCombinerRule(),
	NewGoogleMonitoringCustomServiceInvalidServiceIdRule(),
	NewGoogleMonitoringMetricDescriptorInvalidLaunchStageRule(),
	NewGoogleMonitoringMetricDescriptorInvalidMetricKindRule(),
	NewGoogleMonitoringMetricDescriptorInvalidValueTypeRule(),
	NewGoogleMonitoringSloInvalidCalendarPeriodRule(),
	NewGoogleMonitoringSloInvalidSloIdRule(),
	NewGoogleMonitoringUptimeCheckConfigInvalidCheckerTypeRule(),
	NewGoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule(),
	NewGoogleNotebooksInstanceInvalidBootDiskTypeRule(),
	NewGoogleNotebooksInstanceInvalidDataDiskTypeRule(),
	NewGoogleNotebooksInstanceInvalidDiskEncryptionRule(),
	NewGoogleNotebooksInstanceInvalidNicTypeRule(),
	NewGoogleOsConfigPatchDeploymentInvalidPatchDeploymentIdRule(),
	NewGooglePrivatecaCaPoolInvalidTierRule(),
	NewGooglePrivatecaCertificateAuthorityInvalidTypeRule(),
	NewGooglePubsubSchemaInvalidTypeRule(),
	NewGoogleRedisInstanceInvalidConnectModeRule(),
	NewGoogleRedisInstanceInvalidNameRule(),
	NewGoogleRedisInstanceInvalidReadReplicasModeRule(),
	NewGoogleRedisInstanceInvalidTierRule(),
	NewGoogleRedisInstanceInvalidTransitEncryptionModeRule(),
	NewGoogleSecurityCenterSourceInvalidDisplayNameRule(),
	NewGoogleSpannerDatabaseInvalidDatabaseDialectRule(),
	NewGoogleSpannerDatabaseInvalidNameRule(),
	NewGoogleSpannerInstanceInvalidNameRule(),
	NewGoogleSqlSourceRepresentationInstanceInvalidDatabaseVersionRule(),
	NewGoogleStorageBucketAccessControlInvalidRoleRule(),
	NewGoogleStorageDefaultObjectAccessControlInvalidRoleRule(),
	NewGoogleStorageHmacKeyInvalidStateRule(),
	NewGoogleStorageObjectAccessControlInvalidRoleRule(),
	NewGoogleTagsTagKeyInvalidPurposeRule(),
}
