package repositories

import (
	"cs-ut-ee/build-it-project/pkg/internald/domain"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PurchaseOrderRepository struct {
	gormDB *gorm.DB
}

func NewPurchaseOrderRepository(gormDB *gorm.DB) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{
		gormDB: gormDB,
	}
}

func (por *PurchaseOrderRepository) CreatePO(po *domain.PurchaseOrder) (*domain.PurchaseOrder, error) {
	var res *gorm.DB
	if res = por.gormDB.Create(po); res.Error != nil {
		log.Errorf("Couldn't insert plant hire to db", res.Error)
		return nil, res.Error
	}

	if res.RowsAffected != 1 {
		log.Errorf("Couldn't insert plant hire to db. 0 rows inserted")
		return nil, res.Error
	}

	return po, nil
}

func (por *PurchaseOrderRepository) UpdatePOStatus(oldPo *domain.PurchaseOrder, status string) (*domain.PurchaseOrder, error) {

	var res *gorm.DB
	if res = por.gormDB.Model(oldPo).Update("status", status); res.Error != nil {
		log.Errorf("Error updating purchase order status to sent. Error: %v", res.Error)
		return nil, res.Error
	}

	log.Infof("Updated po %v", oldPo)
	return oldPo, nil
}
