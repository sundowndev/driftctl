package cloudflare

import (
	"os"

	"github.com/snyk/driftctl/enumeration"
	"github.com/snyk/driftctl/enumeration/remote/cloudflare/common"
	"github.com/snyk/driftctl/enumeration/remote/terraform"
	tf "github.com/snyk/driftctl/enumeration/terraform"
)

type CloudflareTerraformProvider struct {
	*terraform.TerraformProvider
	name    string
	version string
}

func NewCloudflareTerraformProvider(version string, progress enumeration.ProgressCounter, configDir string) (*CloudflareTerraformProvider, error) {
	if version == "" {
		version = "3.28.0"
	}
	// Just pass your version and name
	p := &CloudflareTerraformProvider{
		version: version,
		name:    tf.CLOUDFLARE,
	}
	// Use TerraformProviderInstaller to retrieve the provider if needed
	installer, err := tf.NewProviderInstaller(tf.ProviderConfig{
		Key:       p.name,
		Version:   version,
		ConfigDir: configDir,
	})
	if err != nil {
		return nil, err
	}

	tfProvider, err := terraform.NewTerraformProvider(installer, terraform.TerraformProviderConfig{
		Name: p.name,
		GetProviderConfig: func(_ string) interface{} {
			c := p.GetConfig()
			return map[string]interface{}{
				"api_key":              c.ApiKey,
				"email":                c.ApiEmail,
				"api_token":            c.ApiToken,
				"api_user_service_key": c.ApiUserServiceKey,
			}
		},
	}, progress)
	if err != nil {
		return nil, err
	}
	p.TerraformProvider = tfProvider
	return p, err
}

func (p *CloudflareTerraformProvider) GetConfig() common.CloudflareProviderConfig {
	return common.CloudflareProviderConfig{
		ApiKey:            os.Getenv("CLOUDFLARE_API_KEY"),
		ApiEmail:          os.Getenv("CLOUDFLARE_API_EMAIL"),
		ApiToken:          os.Getenv("CLOUDFLARE_API_TOKEN"),
		ApiUserServiceKey: os.Getenv("CLOUDFLARE_API_USER_SERVICE_KEY"),
	}
}

func (p *CloudflareTerraformProvider) Name() string {
	return p.name
}

func (p *CloudflareTerraformProvider) Version() string {
	return p.version
}
