package azurerm

import "github.com/snyk/driftctl/enumeration/resource"

func InitResourcesMetadata(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	initAzureContainerRegistryMetadata(resourceSchemaRepository)
	initAzureFirewallMetadata(resourceSchemaRepository)
	initAzureImageMetaData(resourceSchemaRepository)
	initAzureLoadBalancerMetadata(resourceSchemaRepository)
	initAzurePostgresqlDatabaseMetadata(resourceSchemaRepository)
	initAzurePostgresqlServerMetadata(resourceSchemaRepository)
	initAzurePublicIPMetadata(resourceSchemaRepository)
	initAzureResourceGroupMetadata(resourceSchemaRepository)
	initAzureRouteMetaData(resourceSchemaRepository)
	initAzureRouteTableMetaData(resourceSchemaRepository)
	initAzureVirtualNetworkMetaData(resourceSchemaRepository)
	initAzureNetworkSecurityGroupMetadata(resourceSchemaRepository)
	initAzurePrivateDNSZoneMetaData(resourceSchemaRepository)
	initAzurePrivateDNSARecordMetaData(resourceSchemaRepository)
	initAzurePrivateDNSAAAARecordMetaData(resourceSchemaRepository)
	initAzurePrivateDNSPTRRecordMetaData(resourceSchemaRepository)
	initAzurePrivateDNSSRVRecordMetaData(resourceSchemaRepository)
	initAzurePrivateDNSMXRecordMetaData(resourceSchemaRepository)
	initAzurePrivateDNSTXTRecordMetaData(resourceSchemaRepository)
	initAzureSSHPublicKeyMetaData(resourceSchemaRepository)
	initAzurePrivateDNSCNameRecordMetaData(resourceSchemaRepository)
	initAzureLoadBalancerRuleMetadata(resourceSchemaRepository)
}
