package cloudflare

import (
	remoteerror "github.com/snyk/driftctl/enumeration/remote/error"
	"github.com/snyk/driftctl/enumeration/resource"
	"github.com/snyk/driftctl/enumeration/resource/cloudflare"
)

type CloudflareZoneEnumerator struct {
	repository CloudflareRepository
	factory    resource.ResourceFactory
}

func NewCloudflareZoneEnumerator(repo CloudflareRepository, factory resource.ResourceFactory) *CloudflareZoneEnumerator {
	return &CloudflareZoneEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (g *CloudflareZoneEnumerator) SupportedType() resource.ResourceType {
	return cloudflare.CloudflareZoneResourceType
}

func (g *CloudflareZoneEnumerator) Enumerate() ([]*resource.Resource, error) {
	zones, err := g.repository.ListZones()
	if err != nil {
		return nil, remoteerror.NewResourceListingError(err, string(g.SupportedType()))
	}

	results := make([]*resource.Resource, 0, len(zones))

	for _, zone := range zones {
		results = append(
			results,
			g.factory.CreateAbstractResource(
				string(g.SupportedType()),
				zone.ID,
				map[string]interface{}{},
			),
		)
	}

	return results, err
}
