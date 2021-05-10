package ports

import "cs-ut-ee/build-it-project/pkg/internald/domain"

//Driven port
type PlantHireRepositoryPort interface {
	CreatePlantHire(po *domain.PlantHire) (*domain.PlantHire, error)
	ModifyPlantHire(plantHire *domain.PlantHire, modifiedPlantHire *domain.PlantHire) (*domain.PlantHire, error)
	GetPlantHireById(id int64) (*domain.PlantHire, error)
}

type PurchaseOrderRepositoryPort interface {
	CreatePO(po *domain.PurchaseOrder) (*domain.PurchaseOrder, error)
}

type InvoiceRepositoryPort interface {
	UpdateStatusReturn(id int64) (bool, error)
	UpdateStatusReject(id int64) (bool, error)
}
