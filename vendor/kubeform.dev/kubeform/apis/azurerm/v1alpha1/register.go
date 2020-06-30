/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"kubeform.dev/kubeform/apis/azurerm"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{Group: azurerm.GroupName, Version: "v1alpha1"}

var (
	// TODO: move SchemeBuilder with zz_generated.deepcopy.go to k8s.io/api.
	// localSchemeBuilder and AddToScheme will stay in k8s.io/kubernetes.
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	// We only register manually written functions here. The registration of the
	// generated functions takes place in the generated files. The separation
	// makes the code compile even when the generated files are missing.
	localSchemeBuilder.Register(addKnownTypes)
}

// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,

		&AdvancedThreatProtection{},
		&AdvancedThreatProtectionList{},

		&AnalysisServicesServer{},
		&AnalysisServicesServerList{},

		&ApiManagement{},
		&ApiManagementList{},

		&ApiManagementAPI{},
		&ApiManagementAPIList{},

		&ApiManagementAPIOperation{},
		&ApiManagementAPIOperationList{},

		&ApiManagementAPIOperationPolicy{},
		&ApiManagementAPIOperationPolicyList{},

		&ApiManagementAPIPolicy{},
		&ApiManagementAPIPolicyList{},

		&ApiManagementAPISchema{},
		&ApiManagementAPISchemaList{},

		&ApiManagementAPIVersionSet{},
		&ApiManagementAPIVersionSetList{},

		&ApiManagementAuthorizationServer{},
		&ApiManagementAuthorizationServerList{},

		&ApiManagementBackend{},
		&ApiManagementBackendList{},

		&ApiManagementCertificate{},
		&ApiManagementCertificateList{},

		&ApiManagementDiagnostic{},
		&ApiManagementDiagnosticList{},

		&ApiManagementGroup{},
		&ApiManagementGroupList{},

		&ApiManagementGroupUser{},
		&ApiManagementGroupUserList{},

		&ApiManagementIdentityProviderAad{},
		&ApiManagementIdentityProviderAadList{},

		&ApiManagementIdentityProviderFacebook{},
		&ApiManagementIdentityProviderFacebookList{},

		&ApiManagementIdentityProviderGoogle{},
		&ApiManagementIdentityProviderGoogleList{},

		&ApiManagementIdentityProviderMicrosoft{},
		&ApiManagementIdentityProviderMicrosoftList{},

		&ApiManagementIdentityProviderTwitter{},
		&ApiManagementIdentityProviderTwitterList{},

		&ApiManagementLogger{},
		&ApiManagementLoggerList{},

		&ApiManagementOpenidConnectProvider{},
		&ApiManagementOpenidConnectProviderList{},

		&ApiManagementProduct{},
		&ApiManagementProductList{},

		&ApiManagementProductAPI{},
		&ApiManagementProductAPIList{},

		&ApiManagementProductGroup{},
		&ApiManagementProductGroupList{},

		&ApiManagementProductPolicy{},
		&ApiManagementProductPolicyList{},

		&ApiManagementProperty{},
		&ApiManagementPropertyList{},

		&ApiManagementSubscription{},
		&ApiManagementSubscriptionList{},

		&ApiManagementUser{},
		&ApiManagementUserList{},

		&AppConfiguration{},
		&AppConfigurationList{},

		&AppService{},
		&AppServiceList{},

		&AppServiceActiveSlot{},
		&AppServiceActiveSlotList{},

		&AppServiceCertificate{},
		&AppServiceCertificateList{},

		&AppServiceCertificateOrder{},
		&AppServiceCertificateOrderList{},

		&AppServiceCustomHostnameBinding{},
		&AppServiceCustomHostnameBindingList{},

		&AppServicePlan{},
		&AppServicePlanList{},

		&AppServiceSlot{},
		&AppServiceSlotList{},

		&AppServiceSourceControlToken{},
		&AppServiceSourceControlTokenList{},

		&AppServiceVirtualNetworkSwiftConnection{},
		&AppServiceVirtualNetworkSwiftConnectionList{},

		&ApplicationGateway{},
		&ApplicationGatewayList{},

		&ApplicationInsights{},
		&ApplicationInsightsList{},

		&ApplicationInsightsAPIKey{},
		&ApplicationInsightsAPIKeyList{},

		&ApplicationInsightsAnalyticsItem{},
		&ApplicationInsightsAnalyticsItemList{},

		&ApplicationInsightsWebTest{},
		&ApplicationInsightsWebTestList{},

		&ApplicationSecurityGroup{},
		&ApplicationSecurityGroupList{},

		&AutomationAccount{},
		&AutomationAccountList{},

		&AutomationCertificate{},
		&AutomationCertificateList{},

		&AutomationCredential{},
		&AutomationCredentialList{},

		&AutomationDscConfiguration{},
		&AutomationDscConfigurationList{},

		&AutomationDscNodeconfiguration{},
		&AutomationDscNodeconfigurationList{},

		&AutomationJobSchedule{},
		&AutomationJobScheduleList{},

		&AutomationModule{},
		&AutomationModuleList{},

		&AutomationRunbook{},
		&AutomationRunbookList{},

		&AutomationSchedule{},
		&AutomationScheduleList{},

		&AutomationVariableBool{},
		&AutomationVariableBoolList{},

		&AutomationVariableDatetime{},
		&AutomationVariableDatetimeList{},

		&AutomationVariableInt{},
		&AutomationVariableIntList{},

		&AutomationVariableString{},
		&AutomationVariableStringList{},

		&AutoscaleSetting{},
		&AutoscaleSettingList{},

		&AvailabilitySet{},
		&AvailabilitySetList{},

		&AzureadApplication{},
		&AzureadApplicationList{},

		&AzureadServicePrincipal{},
		&AzureadServicePrincipalList{},

		&AzureadServicePrincipalPassword{},
		&AzureadServicePrincipalPasswordList{},

		&BackupContainerStorageAccount{},
		&BackupContainerStorageAccountList{},

		&BackupPolicyFileShare{},
		&BackupPolicyFileShareList{},

		&BackupPolicyVm{},
		&BackupPolicyVmList{},

		&BackupProtectedFileShare{},
		&BackupProtectedFileShareList{},

		&BackupProtectedVm{},
		&BackupProtectedVmList{},

		&BastionHost{},
		&BastionHostList{},

		&BatchAccount{},
		&BatchAccountList{},

		&BatchApplication{},
		&BatchApplicationList{},

		&BatchCertificate{},
		&BatchCertificateList{},

		&BatchPool{},
		&BatchPoolList{},

		&BotChannelEmail{},
		&BotChannelEmailList{},

		&BotChannelMsTeams{},
		&BotChannelMsTeamsList{},

		&BotChannelSlack{},
		&BotChannelSlackList{},

		&BotChannelsRegistration{},
		&BotChannelsRegistrationList{},

		&BotConnection{},
		&BotConnectionList{},

		&BotWebApp{},
		&BotWebAppList{},

		&CdnEndpoint{},
		&CdnEndpointList{},

		&CdnProfile{},
		&CdnProfileList{},

		&CognitiveAccount{},
		&CognitiveAccountList{},

		&ConnectionMonitor{},
		&ConnectionMonitorList{},

		&ContainerGroup{},
		&ContainerGroupList{},

		&ContainerRegistry{},
		&ContainerRegistryList{},

		&ContainerRegistryWebhook{},
		&ContainerRegistryWebhookList{},

		&ContainerService{},
		&ContainerServiceList{},

		&CosmosdbAccount{},
		&CosmosdbAccountList{},

		&CosmosdbCassandraKeyspace{},
		&CosmosdbCassandraKeyspaceList{},

		&CosmosdbGremlinDatabase{},
		&CosmosdbGremlinDatabaseList{},

		&CosmosdbGremlinGraph{},
		&CosmosdbGremlinGraphList{},

		&CosmosdbMongoCollection{},
		&CosmosdbMongoCollectionList{},

		&CosmosdbMongoDatabase{},
		&CosmosdbMongoDatabaseList{},

		&CosmosdbSQLContainer{},
		&CosmosdbSQLContainerList{},

		&CosmosdbSQLDatabase{},
		&CosmosdbSQLDatabaseList{},

		&CosmosdbTable{},
		&CosmosdbTableList{},

		&Dashboard{},
		&DashboardList{},

		&DataFactory{},
		&DataFactoryList{},

		&DataFactoryDatasetMysql{},
		&DataFactoryDatasetMysqlList{},

		&DataFactoryDatasetPostgresql{},
		&DataFactoryDatasetPostgresqlList{},

		&DataFactoryDatasetSQLServerTable{},
		&DataFactoryDatasetSQLServerTableList{},

		&DataFactoryIntegrationRuntimeManaged{},
		&DataFactoryIntegrationRuntimeManagedList{},

		&DataFactoryLinkedServiceDataLakeStorageGen2{},
		&DataFactoryLinkedServiceDataLakeStorageGen2List{},

		&DataFactoryLinkedServiceMysql{},
		&DataFactoryLinkedServiceMysqlList{},

		&DataFactoryLinkedServicePostgresql{},
		&DataFactoryLinkedServicePostgresqlList{},

		&DataFactoryLinkedServiceSQLServer{},
		&DataFactoryLinkedServiceSQLServerList{},

		&DataFactoryPipeline{},
		&DataFactoryPipelineList{},

		&DataFactoryTriggerSchedule{},
		&DataFactoryTriggerScheduleList{},

		&DataLakeAnalyticsAccount{},
		&DataLakeAnalyticsAccountList{},

		&DataLakeAnalyticsFirewallRule{},
		&DataLakeAnalyticsFirewallRuleList{},

		&DataLakeStore{},
		&DataLakeStoreList{},

		&DataLakeStoreFile{},
		&DataLakeStoreFileList{},

		&DataLakeStoreFirewallRule{},
		&DataLakeStoreFirewallRuleList{},

		&DatabricksWorkspace{},
		&DatabricksWorkspaceList{},

		&DdosProtectionPlan{},
		&DdosProtectionPlanList{},

		&DedicatedHost{},
		&DedicatedHostList{},

		&DedicatedHostGroup{},
		&DedicatedHostGroupList{},

		&DevTestLab{},
		&DevTestLabList{},

		&DevTestLinuxVirtualMachine{},
		&DevTestLinuxVirtualMachineList{},

		&DevTestPolicy{},
		&DevTestPolicyList{},

		&DevTestSchedule{},
		&DevTestScheduleList{},

		&DevTestVirtualNetwork{},
		&DevTestVirtualNetworkList{},

		&DevTestWindowsVirtualMachine{},
		&DevTestWindowsVirtualMachineList{},

		&DevspaceController{},
		&DevspaceControllerList{},

		&DiskEncryptionSet{},
		&DiskEncryptionSetList{},

		&DnsARecord{},
		&DnsARecordList{},

		&DnsAaaaRecord{},
		&DnsAaaaRecordList{},

		&DnsCaaRecord{},
		&DnsCaaRecordList{},

		&DnsCnameRecord{},
		&DnsCnameRecordList{},

		&DnsMxRecord{},
		&DnsMxRecordList{},

		&DnsNsRecord{},
		&DnsNsRecordList{},

		&DnsPtrRecord{},
		&DnsPtrRecordList{},

		&DnsSrvRecord{},
		&DnsSrvRecordList{},

		&DnsTxtRecord{},
		&DnsTxtRecordList{},

		&DnsZone{},
		&DnsZoneList{},

		&EventgridDomain{},
		&EventgridDomainList{},

		&EventgridEventSubscription{},
		&EventgridEventSubscriptionList{},

		&EventgridTopic{},
		&EventgridTopicList{},

		&Eventhub{},
		&EventhubList{},

		&EventhubAuthorizationRule{},
		&EventhubAuthorizationRuleList{},

		&EventhubConsumerGroup{},
		&EventhubConsumerGroupList{},

		&EventhubNamespaceAuthorizationRule{},
		&EventhubNamespaceAuthorizationRuleList{},

		&EventhubNamespaceDisasterRecoveryConfig{},
		&EventhubNamespaceDisasterRecoveryConfigList{},

		&EventhubNamespace_{},
		&EventhubNamespace_List{},

		&ExpressRouteCircuit{},
		&ExpressRouteCircuitList{},

		&ExpressRouteCircuitAuthorization{},
		&ExpressRouteCircuitAuthorizationList{},

		&ExpressRouteCircuitPeering{},
		&ExpressRouteCircuitPeeringList{},

		&Firewall{},
		&FirewallList{},

		&FirewallApplicationRuleCollection{},
		&FirewallApplicationRuleCollectionList{},

		&FirewallNATRuleCollection{},
		&FirewallNATRuleCollectionList{},

		&FirewallNetworkRuleCollection{},
		&FirewallNetworkRuleCollectionList{},

		&Frontdoor{},
		&FrontdoorList{},

		&FrontdoorFirewallPolicy{},
		&FrontdoorFirewallPolicyList{},

		&FunctionApp{},
		&FunctionAppList{},

		&HdinsightHadoopCluster{},
		&HdinsightHadoopClusterList{},

		&HdinsightHbaseCluster{},
		&HdinsightHbaseClusterList{},

		&HdinsightInteractiveQueryCluster{},
		&HdinsightInteractiveQueryClusterList{},

		&HdinsightKafkaCluster{},
		&HdinsightKafkaClusterList{},

		&HdinsightMlServicesCluster{},
		&HdinsightMlServicesClusterList{},

		&HdinsightRserverCluster{},
		&HdinsightRserverClusterList{},

		&HdinsightSparkCluster{},
		&HdinsightSparkClusterList{},

		&HdinsightStormCluster{},
		&HdinsightStormClusterList{},

		&HealthcareService{},
		&HealthcareServiceList{},

		&Image{},
		&ImageList{},

		&IotDps{},
		&IotDpsList{},

		&IotDpsCertificate{},
		&IotDpsCertificateList{},

		&Iothub{},
		&IothubList{},

		&IothubConsumerGroup{},
		&IothubConsumerGroupList{},

		&IothubDps{},
		&IothubDpsList{},

		&IothubDpsCertificate{},
		&IothubDpsCertificateList{},

		&IothubDpsSharedAccessPolicy{},
		&IothubDpsSharedAccessPolicyList{},

		&IothubEndpointEventhub{},
		&IothubEndpointEventhubList{},

		&IothubEndpointServicebusQueue{},
		&IothubEndpointServicebusQueueList{},

		&IothubEndpointServicebusTopic{},
		&IothubEndpointServicebusTopicList{},

		&IothubEndpointStorageContainer{},
		&IothubEndpointStorageContainerList{},

		&IothubFallbackRoute{},
		&IothubFallbackRouteList{},

		&IothubRoute{},
		&IothubRouteList{},

		&IothubSharedAccessPolicy{},
		&IothubSharedAccessPolicyList{},

		&KeyVault{},
		&KeyVaultList{},

		&KeyVaultAccessPolicy{},
		&KeyVaultAccessPolicyList{},

		&KeyVaultCertificate{},
		&KeyVaultCertificateList{},

		&KeyVaultKey{},
		&KeyVaultKeyList{},

		&KeyVaultSecret{},
		&KeyVaultSecretList{},

		&KubernetesCluster{},
		&KubernetesClusterList{},

		&KubernetesClusterNodePool{},
		&KubernetesClusterNodePoolList{},

		&KustoCluster{},
		&KustoClusterList{},

		&KustoDatabase{},
		&KustoDatabaseList{},

		&KustoDatabasePrincipal{},
		&KustoDatabasePrincipalList{},

		&KustoEventhubDataConnection{},
		&KustoEventhubDataConnectionList{},

		&Lb{},
		&LbList{},

		&LbBackendAddressPool{},
		&LbBackendAddressPoolList{},

		&LbNATPool{},
		&LbNATPoolList{},

		&LbNATRule{},
		&LbNATRuleList{},

		&LbOutboundRule{},
		&LbOutboundRuleList{},

		&LbProbe{},
		&LbProbeList{},

		&LbRule{},
		&LbRuleList{},

		&LocalNetworkGateway{},
		&LocalNetworkGatewayList{},

		&LogAnalyticsLinkedService{},
		&LogAnalyticsLinkedServiceList{},

		&LogAnalyticsSolution{},
		&LogAnalyticsSolutionList{},

		&LogAnalyticsWorkspace{},
		&LogAnalyticsWorkspaceList{},

		&LogAnalyticsWorkspaceLinkedService{},
		&LogAnalyticsWorkspaceLinkedServiceList{},

		&LogicAppActionCustom{},
		&LogicAppActionCustomList{},

		&LogicAppActionHTTP{},
		&LogicAppActionHTTPList{},

		&LogicAppTriggerCustom{},
		&LogicAppTriggerCustomList{},

		&LogicAppTriggerHTTPRequest{},
		&LogicAppTriggerHTTPRequestList{},

		&LogicAppTriggerRecurrence{},
		&LogicAppTriggerRecurrenceList{},

		&LogicAppWorkflow{},
		&LogicAppWorkflowList{},

		&ManagedDisk{},
		&ManagedDiskList{},

		&ManagementGroup{},
		&ManagementGroupList{},

		&ManagementLock{},
		&ManagementLockList{},

		&MapsAccount{},
		&MapsAccountList{},

		&MariadbConfiguration{},
		&MariadbConfigurationList{},

		&MariadbDatabase{},
		&MariadbDatabaseList{},

		&MariadbFirewallRule{},
		&MariadbFirewallRuleList{},

		&MariadbServer{},
		&MariadbServerList{},

		&MariadbVirtualNetworkRule{},
		&MariadbVirtualNetworkRuleList{},

		&MarketplaceAgreement{},
		&MarketplaceAgreementList{},

		&MediaServicesAccount{},
		&MediaServicesAccountList{},

		&MetricAlertrule{},
		&MetricAlertruleList{},

		&MonitorActionGroup{},
		&MonitorActionGroupList{},

		&MonitorActivityLogAlert{},
		&MonitorActivityLogAlertList{},

		&MonitorAutoscaleSetting{},
		&MonitorAutoscaleSettingList{},

		&MonitorDiagnosticSetting{},
		&MonitorDiagnosticSettingList{},

		&MonitorLogProfile{},
		&MonitorLogProfileList{},

		&MonitorMetricAlert{},
		&MonitorMetricAlertList{},

		&MonitorMetricAlertrule{},
		&MonitorMetricAlertruleList{},

		&MssqlDatabaseVulnerabilityAssessmentRuleBaseline{},
		&MssqlDatabaseVulnerabilityAssessmentRuleBaselineList{},

		&MssqlElasticpool{},
		&MssqlElasticpoolList{},

		&MssqlServerSecurityAlertPolicy{},
		&MssqlServerSecurityAlertPolicyList{},

		&MssqlServerVulnerabilityAssessment{},
		&MssqlServerVulnerabilityAssessmentList{},

		&MysqlConfiguration{},
		&MysqlConfigurationList{},

		&MysqlDatabase{},
		&MysqlDatabaseList{},

		&MysqlFirewallRule{},
		&MysqlFirewallRuleList{},

		&MysqlServer{},
		&MysqlServerList{},

		&MysqlVirtualNetworkRule{},
		&MysqlVirtualNetworkRuleList{},

		&NatGateway{},
		&NatGatewayList{},

		&NetappAccount{},
		&NetappAccountList{},

		&NetappPool{},
		&NetappPoolList{},

		&NetappSnapshot{},
		&NetappSnapshotList{},

		&NetappVolume{},
		&NetappVolumeList{},

		&NetworkConnectionMonitor{},
		&NetworkConnectionMonitorList{},

		&NetworkDdosProtectionPlan{},
		&NetworkDdosProtectionPlanList{},

		&NetworkInterface{},
		&NetworkInterfaceList{},

		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociation{},
		&NetworkInterfaceApplicationGatewayBackendAddressPoolAssociationList{},

		&NetworkInterfaceApplicationSecurityGroupAssociation{},
		&NetworkInterfaceApplicationSecurityGroupAssociationList{},

		&NetworkInterfaceBackendAddressPoolAssociation{},
		&NetworkInterfaceBackendAddressPoolAssociationList{},

		&NetworkInterfaceNATRuleAssociation{},
		&NetworkInterfaceNATRuleAssociationList{},

		&NetworkPacketCapture{},
		&NetworkPacketCaptureList{},

		&NetworkProfile{},
		&NetworkProfileList{},

		&NetworkSecurityGroup{},
		&NetworkSecurityGroupList{},

		&NetworkSecurityRule{},
		&NetworkSecurityRuleList{},

		&NetworkWatcher{},
		&NetworkWatcherList{},

		&NetworkWatcherFlowLog{},
		&NetworkWatcherFlowLogList{},

		&NotificationHub{},
		&NotificationHubList{},

		&NotificationHubAuthorizationRule{},
		&NotificationHubAuthorizationRuleList{},

		&NotificationHubNamespace_{},
		&NotificationHubNamespace_List{},

		&PacketCapture{},
		&PacketCaptureList{},

		&PointToSiteVPNGateway{},
		&PointToSiteVPNGatewayList{},

		&PolicyAssignment{},
		&PolicyAssignmentList{},

		&PolicyDefinition{},
		&PolicyDefinitionList{},

		&PolicySetDefinition{},
		&PolicySetDefinitionList{},

		&PostgresqlConfiguration{},
		&PostgresqlConfigurationList{},

		&PostgresqlDatabase{},
		&PostgresqlDatabaseList{},

		&PostgresqlFirewallRule{},
		&PostgresqlFirewallRuleList{},

		&PostgresqlServer{},
		&PostgresqlServerList{},

		&PostgresqlVirtualNetworkRule{},
		&PostgresqlVirtualNetworkRuleList{},

		&PrivateDNSARecord{},
		&PrivateDNSARecordList{},

		&PrivateDNSAaaaRecord{},
		&PrivateDNSAaaaRecordList{},

		&PrivateDNSCnameRecord{},
		&PrivateDNSCnameRecordList{},

		&PrivateDNSMxRecord{},
		&PrivateDNSMxRecordList{},

		&PrivateDNSPtrRecord{},
		&PrivateDNSPtrRecordList{},

		&PrivateDNSSrvRecord{},
		&PrivateDNSSrvRecordList{},

		&PrivateDNSZone{},
		&PrivateDNSZoneList{},

		&PrivateDNSZoneVirtualNetworkLink{},
		&PrivateDNSZoneVirtualNetworkLinkList{},

		&PrivateEndpoint{},
		&PrivateEndpointList{},

		&PrivateLinkEndpoint{},
		&PrivateLinkEndpointList{},

		&PrivateLinkService{},
		&PrivateLinkServiceList{},

		&ProximityPlacementGroup{},
		&ProximityPlacementGroupList{},

		&PublicIP{},
		&PublicIPList{},

		&PublicIPPrefix{},
		&PublicIPPrefixList{},

		&RecoveryNetworkMapping{},
		&RecoveryNetworkMappingList{},

		&RecoveryReplicatedVm{},
		&RecoveryReplicatedVmList{},

		&RecoveryServicesFabric{},
		&RecoveryServicesFabricList{},

		&RecoveryServicesProtectedVm{},
		&RecoveryServicesProtectedVmList{},

		&RecoveryServicesProtectionContainer{},
		&RecoveryServicesProtectionContainerList{},

		&RecoveryServicesProtectionContainerMapping{},
		&RecoveryServicesProtectionContainerMappingList{},

		&RecoveryServicesProtectionPolicyVm{},
		&RecoveryServicesProtectionPolicyVmList{},

		&RecoveryServicesReplicationPolicy{},
		&RecoveryServicesReplicationPolicyList{},

		&RecoveryServicesVault{},
		&RecoveryServicesVaultList{},

		&RedisCache{},
		&RedisCacheList{},

		&RedisFirewallRule{},
		&RedisFirewallRuleList{},

		&RelayHybridConnection{},
		&RelayHybridConnectionList{},

		&RelayNamespace{},
		&RelayNamespaceList{},

		&ResourceGroup{},
		&ResourceGroupList{},

		&RoleAssignment{},
		&RoleAssignmentList{},

		&RoleDefinition{},
		&RoleDefinitionList{},

		&Route{},
		&RouteList{},

		&RouteTable{},
		&RouteTableList{},

		&SchedulerJob{},
		&SchedulerJobList{},

		&SchedulerJobCollection{},
		&SchedulerJobCollectionList{},

		&SearchService{},
		&SearchServiceList{},

		&SecurityCenterContact{},
		&SecurityCenterContactList{},

		&SecurityCenterSubscriptionPricing{},
		&SecurityCenterSubscriptionPricingList{},

		&SecurityCenterWorkspace{},
		&SecurityCenterWorkspaceList{},

		&ServiceFabricCluster{},
		&ServiceFabricClusterList{},

		&ServicebusNamespace{},
		&ServicebusNamespaceList{},

		&ServicebusNamespaceAuthorizationRule{},
		&ServicebusNamespaceAuthorizationRuleList{},

		&ServicebusQueue{},
		&ServicebusQueueList{},

		&ServicebusQueueAuthorizationRule{},
		&ServicebusQueueAuthorizationRuleList{},

		&ServicebusSubscription{},
		&ServicebusSubscriptionList{},

		&ServicebusSubscriptionRule{},
		&ServicebusSubscriptionRuleList{},

		&ServicebusTopic{},
		&ServicebusTopicList{},

		&ServicebusTopicAuthorizationRule{},
		&ServicebusTopicAuthorizationRuleList{},

		&SharedImage{},
		&SharedImageList{},

		&SharedImageGallery{},
		&SharedImageGalleryList{},

		&SharedImageVersion{},
		&SharedImageVersionList{},

		&SignalrService{},
		&SignalrServiceList{},

		&SiteRecoveryFabric{},
		&SiteRecoveryFabricList{},

		&SiteRecoveryNetworkMapping{},
		&SiteRecoveryNetworkMappingList{},

		&SiteRecoveryProtectionContainer{},
		&SiteRecoveryProtectionContainerList{},

		&SiteRecoveryProtectionContainerMapping{},
		&SiteRecoveryProtectionContainerMappingList{},

		&SiteRecoveryReplicatedVm{},
		&SiteRecoveryReplicatedVmList{},

		&SiteRecoveryReplicationPolicy{},
		&SiteRecoveryReplicationPolicyList{},

		&Snapshot{},
		&SnapshotList{},

		&SqlActiveDirectoryAdministrator{},
		&SqlActiveDirectoryAdministratorList{},

		&SqlDatabase{},
		&SqlDatabaseList{},

		&SqlElasticpool{},
		&SqlElasticpoolList{},

		&SqlFailoverGroup{},
		&SqlFailoverGroupList{},

		&SqlFirewallRule{},
		&SqlFirewallRuleList{},

		&SqlServer{},
		&SqlServerList{},

		&SqlVirtualNetworkRule{},
		&SqlVirtualNetworkRuleList{},

		&StorageAccount{},
		&StorageAccountList{},

		&StorageAccountNetworkRules{},
		&StorageAccountNetworkRulesList{},

		&StorageBlob{},
		&StorageBlobList{},

		&StorageContainer{},
		&StorageContainerList{},

		&StorageDataLakeGen2Filesystem{},
		&StorageDataLakeGen2FilesystemList{},

		&StorageManagementPolicy{},
		&StorageManagementPolicyList{},

		&StorageQueue{},
		&StorageQueueList{},

		&StorageShare{},
		&StorageShareList{},

		&StorageShareDirectory{},
		&StorageShareDirectoryList{},

		&StorageTable{},
		&StorageTableList{},

		&StorageTableEntity{},
		&StorageTableEntityList{},

		&StreamAnalyticsFunctionJavascriptUdf{},
		&StreamAnalyticsFunctionJavascriptUdfList{},

		&StreamAnalyticsJob{},
		&StreamAnalyticsJobList{},

		&StreamAnalyticsOutputBlob{},
		&StreamAnalyticsOutputBlobList{},

		&StreamAnalyticsOutputEventhub{},
		&StreamAnalyticsOutputEventhubList{},

		&StreamAnalyticsOutputMssql{},
		&StreamAnalyticsOutputMssqlList{},

		&StreamAnalyticsOutputServicebusQueue{},
		&StreamAnalyticsOutputServicebusQueueList{},

		&StreamAnalyticsOutputServicebusTopic{},
		&StreamAnalyticsOutputServicebusTopicList{},

		&StreamAnalyticsReferenceInputBlob{},
		&StreamAnalyticsReferenceInputBlobList{},

		&StreamAnalyticsStreamInputBlob{},
		&StreamAnalyticsStreamInputBlobList{},

		&StreamAnalyticsStreamInputEventhub{},
		&StreamAnalyticsStreamInputEventhubList{},

		&StreamAnalyticsStreamInputIothub{},
		&StreamAnalyticsStreamInputIothubList{},

		&Subnet{},
		&SubnetList{},

		&SubnetNATGatewayAssociation{},
		&SubnetNATGatewayAssociationList{},

		&SubnetNetworkSecurityGroupAssociation{},
		&SubnetNetworkSecurityGroupAssociationList{},

		&SubnetRouteTableAssociation{},
		&SubnetRouteTableAssociationList{},

		&TemplateDeployment{},
		&TemplateDeploymentList{},

		&TrafficManagerEndpoint{},
		&TrafficManagerEndpointList{},

		&TrafficManagerProfile{},
		&TrafficManagerProfileList{},

		&UserAssignedIdentity{},
		&UserAssignedIdentityList{},

		&VirtualHub{},
		&VirtualHubList{},

		&VirtualMachine{},
		&VirtualMachineList{},

		&VirtualMachineDataDiskAttachment{},
		&VirtualMachineDataDiskAttachmentList{},

		&VirtualMachineExtension{},
		&VirtualMachineExtensionList{},

		&VirtualMachineScaleSet{},
		&VirtualMachineScaleSetList{},

		&VirtualNetwork{},
		&VirtualNetworkList{},

		&VirtualNetworkGateway{},
		&VirtualNetworkGatewayList{},

		&VirtualNetworkGatewayConnection{},
		&VirtualNetworkGatewayConnectionList{},

		&VirtualNetworkPeering{},
		&VirtualNetworkPeeringList{},

		&VirtualWAN{},
		&VirtualWANList{},

		&VpnGateway{},
		&VpnGatewayList{},

		&VpnServerConfiguration{},
		&VpnServerConfigurationList{},

		&WebApplicationFirewallPolicy{},
		&WebApplicationFirewallPolicyList{},
	)

	scheme.AddKnownTypes(SchemeGroupVersion,
		&metav1.Status{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
