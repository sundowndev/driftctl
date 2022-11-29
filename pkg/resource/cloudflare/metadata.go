package cloudflare

import "github.com/snyk/driftctl/pkg/resource"

func InitResourcesMetadata(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	initCloudflareZoneMetadata(resourceSchemaRepository)
}
