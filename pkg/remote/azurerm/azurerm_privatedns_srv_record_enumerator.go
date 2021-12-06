package azurerm

import (
	"github.com/snyk/driftctl/pkg/remote/azurerm/repository"
	remoteerror "github.com/snyk/driftctl/pkg/remote/error"
	"github.com/snyk/driftctl/pkg/resource"
	"github.com/snyk/driftctl/pkg/resource/azurerm"
)

type AzurermPrivateDNSSRVRecordEnumerator struct {
	repository repository.PrivateDNSRepository
	factory    resource.ResourceFactory
}

func NewAzurermPrivateDNSSRVRecordEnumerator(repo repository.PrivateDNSRepository, factory resource.ResourceFactory) *AzurermPrivateDNSSRVRecordEnumerator {
	return &AzurermPrivateDNSSRVRecordEnumerator{
		repository: repo,
		factory:    factory,
	}
}

func (e *AzurermPrivateDNSSRVRecordEnumerator) SupportedType() resource.ResourceType {
	return azurerm.AzurePrivateDNSSRVRecordResourceType
}

func (e *AzurermPrivateDNSSRVRecordEnumerator) Enumerate() ([]*resource.Resource, error) {

	zones, err := e.repository.ListAllPrivateZones()
	if err != nil {
		return nil, remoteerror.NewResourceListingErrorWithType(err, string(e.SupportedType()), azurerm.AzurePrivateDNSZoneResourceType)
	}

	results := make([]*resource.Resource, 0)

	for _, zone := range zones {
		records, err := e.repository.ListAllSRVRecords(zone)
		if err != nil {
			return nil, remoteerror.NewResourceListingError(err, string(e.SupportedType()))
		}
		for _, record := range records {
			results = append(
				results,
				e.factory.CreateAbstractResource(
					string(e.SupportedType()),
					*record.ID,
					map[string]interface{}{
						"name":      *record.Name,
						"zone_name": *zone.Name,
					},
				),
			)
		}

	}

	return results, err
}
