package http2

// this file serves the role of the adapter or handler
//it receives an http request and based on it it is able to call the correct service

import (
	"encoding/json"
	"net/http"

	"cs-ut-ee/build-it-project/pkg/internald/domain"
	"cs-ut-ee/build-it-project/pkg/internald/ports"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type HTTPHandler struct {
	plantHireService     ports.PlantHireServicePort
	purchaseOrderService ports.PurchaseOrderServicePort
	invoiceService       ports.InvoiceServicePort
}

func NewHTTPHandler(phs ports.PlantHireServicePort, pos ports.PurchaseOrderServicePort, is ports.InvoiceServicePort) *HTTPHandler {
	return &HTTPHandler{
		plantHireService:     phs,
		purchaseOrderService: pos,
		invoiceService:       is,
	}

}

func (h *HTTPHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/plant-hires", h.CreatePlantHire).Methods(http.MethodPost)

}

func (h *HTTPHandler) CreatePlantHire(w http.ResponseWriter, r *http.Request) {
	var phReq domain.PlantHire
	err := json.NewDecoder(r.Body).Decode(&phReq)
	defer r.Body.Close()

	if err != nil {
		log.Errorf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ph, _ := h.plantHireService.CreatePlantHire(&phReq)

	if ph == nil {
		log.Errorf("Could not create po", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	log.Debug(ph.Id)
	log.Debug(ph)
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(&ph)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

/*
func (h *PlantHandler) ModifyPO(w http.ResponseWriter, r *http.Request) {
	log.Info("Modifying PO dates")
	//here we supose that only date can be changed
	var plantUpdate domain.PlantCheckDTO

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["orderId"], 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	err = json.NewDecoder(r.Body).Decode(&plantUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Infof("New dates: %v", plantUpdate)

	modifiedPo, err := h.plantService.ModifyPO(&plantUpdate, id)

	if err != nil {
		log.Errorf("Couldn't modify %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if modifiedPo == nil {
		log.Errorf("Couldnt update")
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&modifiedPo)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}
*/
