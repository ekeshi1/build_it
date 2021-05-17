package repositories

import (
	"cs-ut-ee/build-it-project/pkg/internald/domain"
	"fmt"

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

func (por *PurchaseOrderRepository) GetAllPurchaseOrders() ([]*domain.PurchaseOrder, error) {
	var purchaseOrders []*domain.PurchaseOrder
	if res := por.gormDB.Find(&purchaseOrders); res.Error != nil {
		log.Errorf("Error retrieving all pos. Error: %v", res.Error)
		return nil, res.Error
	}
	return purchaseOrders, nil

}

func (por *PurchaseOrderRepository) ValidatePurchaseOrderId(id int64) (bool, error) {

	//check if there is a purchase order with this id
	//and check if it is unpaid
	var purchaseOrders []*domain.PurchaseOrder
	if res := por.gormDB.Where(&domain.PurchaseOrder{Id: id}).Not(&domain.PurchaseOrder{Status: domain.InvStatusPaid}).Find(&purchaseOrders); res.Error != nil {
		log.Errorf("Error retrieving all pos. Error: %v", res.Error)
		return false, res.Error
	}

	log.Debugf("Found %v purchase orders with this id !", len(purchaseOrders))
	log.Debugf("%v", &purchaseOrders)
	if len(purchaseOrders) != 1 {
		return false, fmt.Errorf("There should be exactly 1 purchase orders to continue.")

	}

	return true, nil

}

//change this with orm!
func (por *PurchaseOrderRepository) GetPurchaseOrderById(id int64) (*domain.PurchaseOrder, error) {
	var purchaseOrder *domain.PurchaseOrder
	if res := por.gormDB.First(&purchaseOrder, id); res.Error != nil {
		log.Errorf("Error retrieving all pos. Error: %v", res.Error)
		return nil, res.Error
	}
	return purchaseOrder, nil
}

func (por *PurchaseOrderRepository) GetPurchaseOrderByPlantHireId(id int64) (*domain.PurchaseOrder, error) {
	var po *domain.PurchaseOrder
	if res := por.gormDB.First(&po, domain.PurchaseOrder{PlantHireId: id}); res.Error != nil {
		log.Errorf("Couldn't find purchase order by plant id %d. Error %v", id, res.Error)
		return nil, res.Error
	}

	return po, nil
}
