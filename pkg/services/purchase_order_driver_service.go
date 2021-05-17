package services

import (
	"cs-ut-ee/build-it-project/pkg/internald/domain"
)

const (
	SELL_IT_PO_URL = "http://localhost:8080/api/purchase_order"
)

type PurchaseOrderDriverService struct {
}

func NewPurchaseOrderDriverService() *PurchaseOrderDriverService {
	return &PurchaseOrderDriverService{}
}

//make http request to rent it
//this can be changed for another transport layer, only by implementing  PurchaseOrderServiceDriverPort interface.
func (pos *PurchaseOrderDriverService) CreatePurchaseOrder(po *domain.PurchaseOrder) (bool, error) {
	/*
		//create buffer
		payloadBuffer := new(bytes.Buffer)

		json.NewEncoder(payloadBuffer).Encode(po)
		//createdPurchaseOrder
		poResponse, err := http.Post(SELL_IT_PO_URL, "application/json", payloadBuffer)

		log.Debugf("Trying to send created po to supplier( RENT IT)")
		if err != nil {
			log.Errorf("Couldn't post the new purchase order to supplier: ", err)
			return false, err
		}

		if poResponse.StatusCode != http.StatusOK {
			log.Errorf("Post unsucessfull with statuscode: %v", poResponse.Status)
			return false, fmt.Errorf("Failed Http request to supplier with status %v", poResponse.Status)
		}

		log.Debugf("Success")
	*/
	return true, nil

}

func (pos *PurchaseOrderDriverService) ModifyPurchaseOrder(po *domain.PurchaseOrder) (bool, error) {

	return true, nil

}
