package ports

import "cs-ut-ee/build-it-project/pkg/internald/domain"

type PlantHireServicePort interface {
	CreatePlantHire(po *domain.PlantHire) (*domain.PlantHire, error)
	ModifyPlantHire(p []byte, id int64) (*domain.PlantHire, error)
	GetPlantHireById(id int64) (*domain.PlantHire, error)
}

type InvoiceServicePort interface {
	CreateInvoice(inv *domain.Invoice) (*domain.Invoice, error)
	ApproveInvoice(invoiceId int64) error
	GetPurchaseOrderByInvoice(invoiceId int64) (*domain.PurchaseOrder, error)
}
