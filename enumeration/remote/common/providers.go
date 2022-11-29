package common

import (
	tf "github.com/snyk/driftctl/enumeration/terraform"
	"github.com/snyk/driftctl/enumeration/terraform/lock"
)

type RemoteParameter string

const (
	RemoteAWSTerraform        = "aws+tf"
	RemoteGithubTerraform     = "github+tf"
	RemoteGoogleTerraform     = "gcp+tf"
	RemoteAzureTerraform      = "azure+tf"
	RemoteCloudflareTerraform = "cloudflare+tf"
)

var remoteParameterMapping = map[RemoteParameter]*lock.ProviderAddress{
	RemoteAWSTerraform: {
		Hostname:  "registry.terraform.io",
		Namespace: "hashicorp",
		Type:      tf.AWS,
	},
	RemoteGithubTerraform: {
		Hostname:  "registry.terraform.io",
		Namespace: "hashicorp",
		Type:      tf.GITHUB,
	},
	RemoteGoogleTerraform: {
		Hostname:  "registry.terraform.io",
		Namespace: "hashicorp",
		Type:      tf.GOOGLE,
	},
	RemoteAzureTerraform: {
		Hostname:  "registry.terraform.io",
		Namespace: "hashicorp",
		Type:      tf.AZURE,
	},
	RemoteCloudflareTerraform: {
		Hostname:  "registry.terraform.io",
		Namespace: "cloudflare",
		Type:      tf.CLOUDFLARE,
	},
}

func (p RemoteParameter) GetProviderAddress() *lock.ProviderAddress {
	return remoteParameterMapping[p]
}
