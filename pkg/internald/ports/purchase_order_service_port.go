package ports

import "cs-ut-ee/build-it-project/pkg/internald/domain"

//this port is driven by front end requests
type PurchaseOrderServicePort interface {
	CreatePurchaseOrder(po *domain.PurchaseOrder) (*domain.PurchaseOrder, error)
	GetAllPurchaseOrders() ([]*domain.PurchaseOrder, error)
}

//this port is used to drive communication with 3d parties(rent it)
type PurchaseOrderServiceDriverPort interface {
	CreatePurchaseOrder(po *domain.PurchaseOrder) (bool, error)
}
