package services

import (
	"cs-ut-ee/build-it-project/pkg/internald/domain"
	"cs-ut-ee/build-it-project/pkg/internald/ports"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type InvoiceService struct {
	invoiceRepository       ports.InvoiceRepositoryPort
	purchaseOrderRepository ports.PurchaseOrderRepositoryPort
}

func NewInvoiceService(ir ports.InvoiceRepositoryPort, por ports.PurchaseOrderRepositoryPort) *InvoiceService {
	return &InvoiceService{
		invoiceRepository:       ir,
		purchaseOrderRepository: por,
	}
}

func (s *InvoiceService) CreateInvoice(inv *domain.Invoice) (*domain.Invoice, error) {

	//check if a purchase order exist and is unpaid

	isPoUnpaid, err := s.purchaseOrderRepository.ValidatePurchaseOrderId(inv.PurchaseOrderId)

	if err != nil || isPoUnpaid == false {
		log.Errorf("Purchase Order Id not validated. Can't accept invoice. %v ", err)
		return nil, err
	}

	log.Infof("Purchase order Id validated. Saving invoice in db.")
	inv.PaymentStatus = domain.InvStatusCreated
	createdInvoice, err := s.invoiceRepository.CreateInvoice(inv)

	if err != nil {
		log.Errorf("Couldn't create invoice with error: ", err)
		return nil, err
	}

	log.Debugf("Created invoice with id : ", createdInvoice.Id)
	return createdInvoice, nil
}

func (s *InvoiceService) ApproveInvoice(id int64) error {
	//here may need to check if status is not paid
	inv, err := s.invoiceRepository.GetInvoice(id)

	if err != nil {
		return err
	}

	if inv.PaymentStatus != domain.InvStatusCreated {
		return fmt.Errorf("Can't approve this invoice in this status!s")
	}

	if err = s.invoiceRepository.UpdateStatus(id, domain.InvStatusApproved); err != nil {
		return err
	}

	return nil
}

func (s *InvoiceService) GetPurchaseOrderByInvoice(id int64) (*domain.PurchaseOrder, error) {
	inv, err := s.invoiceRepository.GetInvoice(id)

	if err != nil {
		return nil, err
	}

	po, err := s.purchaseOrderRepository.GetPurchaseOrderById(inv.PurchaseOrderId)

	if err != nil {
		log.Errorf("Couldn't get purchase order with error: ", err)
		return nil, err
	}

	return po, nil
}
