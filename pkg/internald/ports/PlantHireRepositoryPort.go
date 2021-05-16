package ports

import "cs-ut-ee/build-it-project/pkg/internald/domain"

//Driven port
type PlantHireRepositoryPort interface {
	CreatePlantHire(po *domain.PlantHire) (*domain.PlantHire, error)
	ModifyPlantHire(plantHire *domain.PlantHire, modifiedPlantHire *domain.PlantHire) (*domain.PlantHire, error)
	GetPlantHireById(id int64) (*domain.PlantHire, error)
	CalculatePrice(start string, end string, pricePerDay float64) (float64, error)
}

type PurchaseOrderRepositoryPort interface {
	CreatePO(po *domain.PurchaseOrder) (*domain.PurchaseOrder, error)
	UpdatePOStatus(oldPo *domain.PurchaseOrder, status string) (*domain.PurchaseOrder, error)
	GetAllPurchaseOrders() ([]*domain.PurchaseOrder, error)
	ValidatePurchaseOrderId(id int64) (bool, error)
	GetPurchaseOrderById(id int64) (*domain.PurchaseOrder, error)
}

type InvoiceRepositoryPort interface {
	//UpdateStatusReturn(id int64) (bool, error)
	//UpdateStatusReject(id int64) (bool, error)
	CreateInvoice(inv *domain.Invoice) (*domain.Invoice, error)
	UpdateStatus(invoiceId int64, newStatus string) error
	GetInvoice(id int64) (*domain.Invoice, error)
}
