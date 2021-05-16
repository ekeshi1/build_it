package http2

// this file serves the role of the adapter or handler
//it receives an http request and based on it it is able to call the correct service

import (
	"cs-ut-ee/build-it-project/pkg/internald/domain"
	"cs-ut-ee/build-it-project/pkg/internald/ports"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"strings"

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
	router.HandleFunc("/api/plant-hires/{id}", h.ModifyPlantHireDates).Methods(http.MethodPatch)
	router.HandleFunc("/api/plant-hires/{id}", h.GetPlantHireById).Methods(http.MethodGet)
	router.HandleFunc("/api/plant-hires/{id}/status", h.ModifyPlantHireStatus).Methods(http.MethodPatch)
	//router.HandleFunc("/api/plant-hires/{id}/purchase-order", h.createPO).Methods(http.MethodPost)
	router.HandleFunc("/api/plant-hires/{id}/extension", h.ModifyPlantHireExtension).Methods(http.MethodPut)

}

func (h *HTTPHandler) RegisterPORoutes(router *mux.Router) {
	router.HandleFunc("/api/purchase-orders", h.GetAllPurchaseOrders).Methods(http.MethodGet)
}

func (h *HTTPHandler) RegisterInvoiceRoutes(router *mux.Router) {
	router.HandleFunc("/api/invoices/{id}", h.GetInvoiceById).Methods(http.MethodGet)
	router.HandleFunc("/api/invoices", h.SubmitInvoice).Methods(http.MethodPost)
	router.HandleFunc("/api/invoices/{invoiceId}/approve", h.ApproveInvoice).Methods(http.MethodPost)
	router.HandleFunc("/api/invoices/{invoiceId}/purchase-order", h.GetInvoicePo).Methods(http.MethodGet)
}

func (h *HTTPHandler) GetInvoiceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.ParseInt(vars["id"], 10, 64)
	plants, err := h.invoiceService.GetInvoice(key)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, "Invoice with this id does not exist", http.StatusNotFound)
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&plants)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (h *HTTPHandler) GetInvoicePo(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	invoiceId, err := strconv.ParseInt(vars["invoiceId"], 10, 64)

	if err != nil {
		http.Error(w, "Bad URL", http.StatusBadRequest)
		return
	}

	po, err := h.invoiceService.GetPurchaseOrderByInvoice(invoiceId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&po)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (h *HTTPHandler) ApproveInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invoiceId, err := strconv.ParseInt(vars["invoiceId"], 10, 64)
	if err != nil {
		http.Error(w, "Bad URL", http.StatusBadRequest)
		return
	}
	err = h.invoiceService.ApproveInvoice(invoiceId)

	log.Error(err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err1 := h.invoiceService.PayInvoice(invoiceId)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode("Invoice is approved, paid and the third party is informed.")
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (h *HTTPHandler) SubmitInvoice(w http.ResponseWriter, r *http.Request) {
	var inv domain.Invoice
	err := json.NewDecoder(r.Body).Decode(&inv)
	log.Info(inv)
	defer r.Body.Close()
	if err != nil {
		log.Errorf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdInvoice, err := h.invoiceService.CreateInvoice(&inv)

	if createdInvoice == nil || err != nil {
		log.Errorf("Could not create invoice", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}
}

func (h *HTTPHandler) GetAllPurchaseOrders(w http.ResponseWriter, r *http.Request) {

	pos, err := h.purchaseOrderService.GetAllPurchaseOrders()

	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&pos)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}
func (h *HTTPHandler) createPO(w http.ResponseWriter, r *http.Request) {
	var poReq domain.PurchaseOrder
	err := json.NewDecoder(r.Body).Decode(&poReq)
	defer r.Body.Close()
	if err != nil {
		log.Errorf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	po, _ := h.purchaseOrderService.CreatePurchaseOrder(&poReq)

	if po == nil {
		log.Errorf("Could not create po", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

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
		log.Errorf("Could not create plant hire", err)
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

func (h *HTTPHandler) ModifyPlantHireDates(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.ParseInt(vars["id"], 10, 64)

	patchJSON, _ := ioutil.ReadAll(r.Body)

	p1 := strings.Contains(string(patchJSON), "plantArrivalDate")
	p2 := strings.Contains(string(patchJSON), "plantReturnDate")
	if p1 == false && p2 == false {
		http.Error(w, "It is now allowed to update this data", http.StatusBadRequest)
		return
	}

	mph, _ := h.plantHireService.ModifyPlantHire(patchJSON, key)

	if mph == nil {
		log.Errorf("Could not modify ph", err)
		http.Error(w, "Plant with this id does not exist", http.StatusNotFound)
		return
	}

	log.Debug(mph.Id)
	log.Debug(mph)
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&mph)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *HTTPHandler) GetPlantHireById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.ParseInt(vars["id"], 10, 64)
	plants, err := h.plantHireService.GetPlantHireById(key)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, "Plant with this id does not exist", http.StatusNotFound)
		return
	}
	// write success response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&plants)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
	}
}

func (h *HTTPHandler) ModifyPlantHireStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.ParseInt(vars["id"], 10, 64)

	patchJSON, _ := ioutil.ReadAll(r.Body)

	p1 := strings.Contains(string(patchJSON), "status")
	b1 := strings.Contains(string(patchJSON), "APPROVED")
	b2 := strings.Contains(string(patchJSON), "REJECTED")
	b3 := strings.Contains(string(patchJSON), "MODIFIED")

	s1 := strings.Contains(string(patchJSON), "approved")
	s2 := strings.Contains(string(patchJSON), "rejected")
	s3 := strings.Contains(string(patchJSON), "modified")

	if p1 == false {
		http.Error(w, "It is now allowed to update this data", http.StatusBadRequest)
		return
	}

	if b1 == false && b2 == false && b3 == false {
		if s1 == true || s2 == true || s3 == true {
			http.Error(w, "The value of status should be uppercase", http.StatusBadRequest)
			return
		}
		http.Error(w, "The value is not correct", http.StatusBadRequest)
		return
	}

	mph, _ := h.plantHireService.ModifyPlantHire(patchJSON, key)

	if mph == nil {
		log.Errorf("Could not modify ph", err)
		http.Error(w, "Plant with this id does not exist", http.StatusNotFound)
		return
	}

	if mph.Status == "APPROVED" {
		var poReq domain.PurchaseOrder
		poReq.PlantHireId = mph.Id
		poReq.Description = "Purchase order is created"
		poReq.Creator = "BUILD_IT"

		po, _ := h.purchaseOrderService.CreatePurchaseOrder(&poReq)

		if po == nil {
			log.Errorf("Could not create po", err)
			http.Error(w, "Could not create po but status updated", http.StatusBadRequest)
			return
		}
	}

	log.Debug(mph.Id)
	log.Debug(mph)
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&mph)
	if err != nil {
		log.Errorf("Could not encode json, err %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *HTTPHandler) ModifyPlantHireExtension(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		log.Errorf("Error while parsing id, err %v", err)
	}

	plantHireExtensionDTO := &domain.PlantHireExtensionDTO{}
	err1 := json.NewDecoder(r.Body).Decode(plantHireExtensionDTO)
	defer r.Body.Close()

	if err1 != nil {
		log.Errorf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mph, _ := h.plantHireService.ModifyPlantHireExtension(key, plantHireExtensionDTO)

	if mph == nil {
		log.Errorf("Could not modify plant hire", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	log.Debug(mph.Id)
	log.Debug(mph)
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&mph)
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
