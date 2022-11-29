package cloudflare

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/pkg/errors"
	"github.com/snyk/driftctl/enumeration"
	"github.com/snyk/driftctl/enumeration/alerter"
	"github.com/snyk/driftctl/enumeration/remote/cache"
	commonCF "github.com/snyk/driftctl/enumeration/remote/cloudflare/common"
	"github.com/snyk/driftctl/enumeration/remote/common"
	"github.com/snyk/driftctl/enumeration/resource"
	"github.com/snyk/driftctl/enumeration/terraform"
)

func Init(version string, alerter alerter.AlerterInterface, providerLibrary *terraform.ProviderLibrary, remoteLibrary *common.RemoteLibrary, progress enumeration.ProgressCounter, factory resource.ResourceFactory, configDir string) error {

	provider, err := NewCloudflareTerraformProvider(version, progress, configDir)
	if err != nil {
		return err
	}
	err = provider.Init()
	if err != nil {
		return err
	}

	providerConfig := provider.GetConfig()

	client, err := NewClient(providerConfig)
	if err != nil {
		return err
	}

	repo := NewCloudflareRepository(client, cache.New(100))

	providerLibrary.AddProvider(terraform.CLOUDFLARE, provider)

	remoteLibrary.AddEnumerator(NewCloudflareZoneEnumerator(repo, factory))

	return nil
}

func NewClient(providerConfig commonCF.CloudflareProviderConfig) (*cloudflare.API, error) {
	if providerConfig.ApiKey != "" || providerConfig.ApiEmail != "" {
		api, err := cloudflare.New(providerConfig.ApiKey, providerConfig.ApiEmail)
		return api, err
	}
	if providerConfig.ApiToken != "" {
		api, err := cloudflare.NewWithAPIToken(providerConfig.ApiToken)
		return api, err
	}
	if providerConfig.ApiUserServiceKey != "" {
		api, err := cloudflare.NewWithUserServiceKey(providerConfig.ApiUserServiceKey)
		return api, err
	}
	return nil, errors.New("couldn't find a way to authenticate to Cloudflare API")
}
