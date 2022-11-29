package cloudflare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/snyk/driftctl/enumeration/remote/cache"
)

type CloudflareRepository interface {
	ListZones() ([]cloudflare.Zone, error)
}

type cloudflareRepository struct {
	client *cloudflare.API
	cache  cache.Cache
}

func NewCloudflareRepository(client *cloudflare.API, c cache.Cache) *cloudflareRepository {
	return &cloudflareRepository{
		client: client,
		cache:  c,
	}
}

func (r *cloudflareRepository) ListZones() ([]cloudflare.Zone, error) {
	return r.client.ListZones(context.Background())
}
