package repositories

import (
	"cs-ut-ee/build-it-project/pkg/internald/domain"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type InvoiceRepository struct {
	gormDB *gorm.DB
}

func NewInvoiceRepository(gormDB *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{
		gormDB: gormDB,
	}
}

func (phr *InvoiceRepository) CreateInvoice(inv *domain.Invoice) (*domain.Invoice, error) {
	var res *gorm.DB
	if res = phr.gormDB.Create(inv); res.Error != nil {
		log.Errorf("Couldn't insert invoice to db", res.Error)
		return nil, res.Error
	}

	if res.RowsAffected != 1 {
		log.Errorf("Couldn't insert plant hire to db. 0 rows inserted")
		return nil, res.Error
	}

	return inv, nil
}

func (phr *InvoiceRepository) GetInvoice(id int64) (*domain.Invoice, error) {
	var inv *domain.Invoice
	if res := phr.gormDB.First(&inv, domain.Invoice{Id: id}); res.Error != nil {
		log.Errorf("Couldn't find invoice with id %d. Error %v", id, res.Error)
		return nil, res.Error
	}

	return inv, nil
}

func (phr *InvoiceRepository) UpdateStatus(id int64, newStatus string) error {
	var inv *domain.Invoice
	if res := phr.gormDB.First(&inv, &domain.Invoice{Id: id}); res.Error != nil {
		log.Errorf("Couldn't find invoice with id %d. Error %v", id, res.Error)
		return res.Error
	}
	inv.PaymentStatus = newStatus
	if res := phr.gormDB.Save(&inv); res.Error != nil {
		log.Errorf("Couldn't update status for invoice with id %d. Error %v", id, res.Error)
		return res.Error
	}
	return nil
}
