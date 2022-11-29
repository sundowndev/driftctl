package cloudflare

import (
	"github.com/snyk/driftctl/enumeration/resource"
	dctlresource "github.com/snyk/driftctl/pkg/resource"
)

const CloudflareZoneResourceType = "cloudflare_zone"

func initCloudflareZoneMetadata(resourceSchemaRepository dctlresource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetHumanReadableAttributesFunc(CloudflareZoneResourceType, func(res *resource.Resource) map[string]string {
		attrs := make(map[string]string)
		if name := res.Attrs.GetString("zone"); name != nil && *name != "" {
			attrs["Name"] = *name
		}
		return attrs
	})
}
