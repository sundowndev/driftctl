package aws

import "github.com/snyk/driftctl/pkg/resource"

const AwsNatGatewayResourceType = "aws_nat_gateway"

func initNatGatewayMetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetFlags(AwsNatGatewayResourceType, resource.FlagDeepMode)
}
