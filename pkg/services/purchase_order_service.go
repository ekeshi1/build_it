package services

import (
	"cs-ut-ee/build-it-project/pkg/internald/domain"
	"cs-ut-ee/build-it-project/pkg/internald/ports"

	log "github.com/sirupsen/logrus"
)

type PurchaseOrderService struct {
	purchaseOrderRepository ports.PurchaseOrderRepositoryPort
	poDriverPort            ports.PurchaseOrderServiceDriverPort
}

func NewPurchaseOrderService(pos ports.PurchaseOrderRepositoryPort, poDriver ports.PurchaseOrderServiceDriverPort) *PurchaseOrderService {
	return &PurchaseOrderService{
		purchaseOrderRepository: pos,
		poDriverPort:            poDriver,
	}
}

func (pos *PurchaseOrderService) CreatePurchaseOrder(po *domain.PurchaseOrder) (*domain.PurchaseOrder, error) {
	//createdPurchaseOrder
	po.Status = domain.POStatusCreated
	po.DeliveryStatus = domain.PODeliveryStatusCreated
	createdPurchaseOrder, err := pos.purchaseOrderRepository.CreatePO(po)

	if err != nil {
		log.Errorf("Couldn't create new plant with error: ", err)
		return nil, err
	}

	log.Debugf("Create purchase order with id : ", createdPurchaseOrder.Id)

	log.Debugf("Trying to send created po to supplier( RENT IT)")

	//make http request to rent it
	//this can be changed for another transport layer, only by implementing  PurchaseOrderServiceDriverPort interface.
	if isSuccessfull, err := pos.poDriverPort.CreatePurchaseOrder(createdPurchaseOrder); err != nil || isSuccessfull != true {
		log.Error("Somethihng went wrong notifying third party about po. Error %v", err)
		return nil, err
	}

	updatedPo, err := pos.purchaseOrderRepository.UpdatePOStatus(createdPurchaseOrder, domain.POStatusSent)

	if err != nil {
		return nil, err
	}
	log.Debugf("Updated PO status to sent")

	return updatedPo, nil
}

func (pos *PurchaseOrderService) GetAllPurchaseOrders() ([]*domain.PurchaseOrder, error) {

	pOrders, err := pos.purchaseOrderRepository.GetAllPurchaseOrders()

	if err != nil {
		log.Errorf("Couldn't get pos. Err : ", err)
		return nil, err
	}

	return pOrders, nil

}

func (pos *PurchaseOrderService) GetPurchaseOrderByPlantHireId(id int64) (*domain.PurchaseOrder, error) {
	po, err := pos.purchaseOrderRepository.GetPurchaseOrderByPlantHireId(id)
	if err != nil {
		log.Errorf("Couldn't get purchase order with error: ", err)
		return nil, err
	}
	return po, nil
}
