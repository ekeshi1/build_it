package ports

import "cs-ut-ee/build-it-project/pkg/internald/domain"

type PlantHireServicePort interface {
	CreatePlantHire(po *domain.PlantHire) (*domain.PlantHire, error)
	ModifyPlantHire(p []byte, id int64) (*domain.PlantHire, error)
	GetPlantHireById(id int64) (*domain.PlantHire, error)
}

type InvoiceServicePort interface {
}
