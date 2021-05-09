package ports

import "cs-ut-ee/build-it-project/pkg/internald/domain"

type PlantHireServicePort interface {
	CreatePlantHire(po *domain.PlantHire) (*domain.PlantHire, error)
}

type PurchaseOrderServicePort interface {
}

type InvoiceServicePort interface {
}
