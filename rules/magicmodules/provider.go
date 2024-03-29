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
	NewGoogleAlloydbBackupInvalidTypeRule(),
	NewGoogleAlloydbClusterInvalidClusterTypeRule(),
	NewGoogleAlloydbInstanceInvalidAvailabilityTypeRule(),
	NewGoogleAlloydbInstanceInvalidInstanceTypeRule(),
	NewGoogleAlloydbUserInvalidUserTypeRule(),
	NewGoogleApigeeEnvironmentInvalidApiProxyTypeRule(),
	NewGoogleApigeeEnvironmentInvalidDeploymentTypeRule(),
	NewGoogleApigeeEnvironmentInvalidTypeRule(),
	NewGoogleApigeeOrganizationInvalidRetentionRule(),
	NewGoogleApigeeOrganizationInvalidRuntimeTypeRule(),
	NewGoogleApigeeTargetServerInvalidProtocolRule(),
	NewGoogleAppEngineDomainMappingInvalidOverrideStrategyRule(),
	NewGoogleAppEngineFirewallRuleInvalidActionRule(),
	NewGoogleAppEngineFlexibleAppVersionInvalidServingStatusRule(),
	NewGoogleArtifactRegistryRepositoryInvalidModeRule(),
	NewGoogleBeyondcorpAppGatewayInvalidHostTypeRule(),
	NewGoogleBeyondcorpAppGatewayInvalidTypeRule(),
	NewGoogleBiglakeTableInvalidTypeRule(),
	NewGoogleBigqueryDatapolicyDataPolicyInvalidDataPolicyTypeRule(),
	NewGoogleBigqueryRoutineInvalidDeterminismLevelRule(),
	NewGoogleBigqueryRoutineInvalidLanguageRule(),
	NewGoogleBigqueryRoutineInvalidRoutineTypeRule(),
	NewGoogleBinaryAuthorizationPolicyInvalidGlobalPolicyEvaluationModeRule(),
	NewGoogleBlockchainNodeEngineBlockchainNodesInvalidBlockchainTypeRule(),
	NewGoogleCertificateManagerCertificateIssuanceConfigInvalidKeyAlgorithmRule(),
	NewGoogleCloudAssetFolderFeedInvalidContentTypeRule(),
	NewGoogleCloudAssetOrganizationFeedInvalidContentTypeRule(),
	NewGoogleCloudAssetProjectFeedInvalidContentTypeRule(),
	NewGoogleCloudIdentityGroupInvalidInitialGroupConfigRule(),
	NewGoogleCloudIdsEndpointInvalidSeverityRule(),
	NewGoogleCloudRunV2JobInvalidLaunchStageRule(),
	NewGoogleCloudRunV2ServiceInvalidIngressRule(),
	NewGoogleCloudRunV2ServiceInvalidLaunchStageRule(),
	NewGoogleCloudbuildTriggerInvalidIncludeBuildLogsRule(),
	NewGoogleComputeAddressInvalidAddressTypeRule(),
	NewGoogleComputeAddressInvalidIpVersionRule(),
	NewGoogleComputeAddressInvalidIpv6EndpointTypeRule(),
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
	NewGoogleComputeForwardingRuleInvalidIpVersionRule(),
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
	NewGoogleComputeInterconnectAttachmentInvalidStackTypeRule(),
	NewGoogleComputeInterconnectAttachmentInvalidTypeRule(),
	NewGoogleComputeManagedSslCertificateInvalidTypeRule(),
	NewGoogleComputeNetworkEndpointGroupInvalidNetworkEndpointTypeRule(),
	NewGoogleComputeNetworkInvalidNetworkFirewallPolicyEnforcementOrderRule(),
	NewGoogleComputeNodeTemplateInvalidCpuOvercommitTypeRule(),
	NewGoogleComputeRegionBackendServiceInvalidLoadBalancingSchemeRule(),
	NewGoogleComputeRegionBackendServiceInvalidLocalityLbPolicyRule(),
	NewGoogleComputeRegionBackendServiceInvalidProtocolRule(),
	NewGoogleComputeRegionBackendServiceInvalidSessionAffinityRule(),
	NewGoogleComputeRegionCommitmentInvalidCategoryRule(),
	NewGoogleComputeRegionCommitmentInvalidPlanRule(),
	NewGoogleComputeRegionNetworkEndpointGroupInvalidNetworkEndpointTypeRule(),
	NewGoogleComputeRegionSslPolicyInvalidMinTlsVersionRule(),
	NewGoogleComputeRegionSslPolicyInvalidProfileRule(),
	NewGoogleComputeRegionTargetTcpProxyInvalidProxyHeaderRule(),
	NewGoogleComputeRouteInvalidNameRule(),
	NewGoogleComputeRouterNatInvalidNatIpAllocateOptionRule(),
	NewGoogleComputeRouterNatInvalidSourceSubnetworkIpRangesToNatRule(),
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
	NewGoogleDataPipelinePipelineInvalidStateRule(),
	NewGoogleDataPipelinePipelineInvalidTypeRule(),
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
	NewGoogleDialogflowCxSecuritySettingsInvalidRedactionScopeRule(),
	NewGoogleDialogflowCxSecuritySettingsInvalidRedactionStrategyRule(),
	NewGoogleDialogflowCxSecuritySettingsInvalidRetentionStrategyRule(),
	NewGoogleDialogflowEntityTypeInvalidKindRule(),
	NewGoogleDialogflowIntentInvalidWebhookStateRule(),
	NewGoogleDiscoveryEngineChatEngineInvalidIndustryVerticalRule(),
	NewGoogleDiscoveryEngineDataStoreInvalidContentConfigRule(),
	NewGoogleDiscoveryEngineDataStoreInvalidIndustryVerticalRule(),
	NewGoogleDiscoveryEngineSearchEngineInvalidIndustryVerticalRule(),
	NewGoogleDnsManagedZoneInvalidVisibilityRule(),
	NewGoogleDocumentAiWarehouseLocationInvalidAccessControlModeRule(),
	NewGoogleDocumentAiWarehouseLocationInvalidDatabaseTypeRule(),
	NewGoogleDocumentAiWarehouseLocationInvalidDocumentCreatorDefaultRoleRule(),
	NewGoogleEdgecontainerClusterInvalidReleaseChannelRule(),
	NewGoogleFirebaseAppCheckServiceConfigInvalidEnforcementModeRule(),
	NewGoogleFirestoreDatabaseInvalidAppEngineIntegrationModeRule(),
	NewGoogleFirestoreDatabaseInvalidConcurrencyModeRule(),
	NewGoogleFirestoreDatabaseInvalidDeleteProtectionStateRule(),
	NewGoogleFirestoreDatabaseInvalidPointInTimeRecoveryEnablementRule(),
	NewGoogleFirestoreDatabaseInvalidTypeRule(),
	NewGoogleFirestoreIndexInvalidApiScopeRule(),
	NewGoogleFirestoreIndexInvalidQueryScopeRule(),
	NewGoogleHealthcareFhirStoreInvalidComplexDataTypeReferenceParsingRule(),
	NewGoogleHealthcareFhirStoreInvalidVersionRule(),
	NewGoogleIntegrationConnectorsConnectionInvalidEventingEnablementTypeRule(),
	NewGoogleKmsCryptoKeyVersionInvalidStateRule(),
	NewGoogleKmsKeyRingImportJobInvalidImportMethodRule(),
	NewGoogleKmsKeyRingImportJobInvalidProtectionLevelRule(),
	NewGoogleLookerInstanceInvalidNameRule(),
	NewGoogleLookerInstanceInvalidPlatformEditionRule(),
	NewGoogleMemcacheInstanceInvalidMemcacheVersionRule(),
	NewGoogleMonitoringAlertPolicyInvalidCombinerRule(),
	NewGoogleMonitoringAlertPolicyInvalidSeverityRule(),
	NewGoogleMonitoringCustomServiceInvalidServiceIdRule(),
	NewGoogleMonitoringMetricDescriptorInvalidLaunchStageRule(),
	NewGoogleMonitoringMetricDescriptorInvalidMetricKindRule(),
	NewGoogleMonitoringMetricDescriptorInvalidValueTypeRule(),
	NewGoogleMonitoringSloInvalidCalendarPeriodRule(),
	NewGoogleMonitoringSloInvalidSloIdRule(),
	NewGoogleMonitoringUptimeCheckConfigInvalidCheckerTypeRule(),
	NewGoogleNetappStoragePoolInvalidServiceLevelRule(),
	NewGoogleNetappVolumeInvalidSecurityStyleRule(),
	NewGoogleNetworkConnectivityPolicyBasedRouteInvalidNextHopOtherRoutesRule(),
	NewGoogleNetworkSecurityAddressGroupInvalidTypeRule(),
	NewGoogleNetworkSecurityGatewaySecurityPolicyRuleInvalidBasicProfileRule(),
	NewGoogleNetworkServicesEdgeCacheOriginInvalidProtocolRule(),
	NewGoogleNetworkServicesGatewayInvalidTypeRule(),
	NewGoogleNotebooksInstanceInvalidBootDiskTypeRule(),
	NewGoogleNotebooksInstanceInvalidDataDiskTypeRule(),
	NewGoogleNotebooksInstanceInvalidDiskEncryptionRule(),
	NewGoogleNotebooksInstanceInvalidNicTypeRule(),
	NewGoogleOrgPolicyCustomConstraintInvalidActionTypeRule(),
	NewGoogleOsConfigPatchDeploymentInvalidPatchDeploymentIdRule(),
	NewGooglePrivatecaCaPoolInvalidTierRule(),
	NewGooglePrivatecaCertificateAuthorityInvalidTypeRule(),
	NewGooglePubsubSchemaInvalidTypeRule(),
	NewGoogleRedisClusterInvalidAuthorizationModeRule(),
	NewGoogleRedisClusterInvalidTransitEncryptionModeRule(),
	NewGoogleRedisInstanceInvalidConnectModeRule(),
	NewGoogleRedisInstanceInvalidNameRule(),
	NewGoogleRedisInstanceInvalidReadReplicasModeRule(),
	NewGoogleRedisInstanceInvalidTierRule(),
	NewGoogleRedisInstanceInvalidTransitEncryptionModeRule(),
	NewGoogleSccEventThreatDetectionCustomModuleInvalidEnablementStateRule(),
	NewGoogleSccFolderCustomModuleInvalidEnablementStateRule(),
	NewGoogleSccOrganizationCustomModuleInvalidEnablementStateRule(),
	NewGoogleSccProjectCustomModuleInvalidEnablementStateRule(),
	NewGoogleSccSourceInvalidDisplayNameRule(),
	NewGoogleSecurityposturePostureInvalidStateRule(),
	NewGoogleSpannerDatabaseInvalidDatabaseDialectRule(),
	NewGoogleSpannerDatabaseInvalidNameRule(),
	NewGoogleSpannerInstanceInvalidNameRule(),
	NewGoogleSqlSourceRepresentationInstanceInvalidDatabaseVersionRule(),
	NewGoogleStorageBucketAccessControlInvalidRoleRule(),
	NewGoogleStorageDefaultObjectAccessControlInvalidRoleRule(),
	NewGoogleStorageHmacKeyInvalidStateRule(),
	NewGoogleStorageObjectAccessControlInvalidRoleRule(),
	NewGoogleTagsTagKeyInvalidPurposeRule(),
	NewGoogleVmwareengineExternalAccessRuleInvalidActionRule(),
	NewGoogleVmwareengineNetworkInvalidTypeRule(),
	NewGoogleVmwareengineNetworkPeeringInvalidPeerNetworkTypeRule(),
	NewGoogleVmwareenginePrivateCloudInvalidTypeRule(),
	NewGoogleWorkflowsWorkflowInvalidCallLogLevelRule(),
}
