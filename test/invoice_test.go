package test

import (
	"bytes"
	"cs-ut-ee/build-it-project/pkg/internald/domain"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

//CC9
func TestSubmitInvoice(t *testing.T) {
	url1 := "http://localhost:8081/api/invoices"

	url2 := "http://localhost:8081/api/plant-hires/2/status"

	var jsonStr1 = []byte(`[{"op": "replace", "path": "/status","value": "APPROVED"}]`)
	req1, _ := http.NewRequest("PATCH", url2, bytes.NewBuffer(jsonStr1))

	client1 := &http.Client{}
	resp1, err1 := client1.Do(req1)
	if err1 != nil {
		panic(err1)
	}
	defer resp1.Body.Close()

	var jsonStr2 = []byte(`{"PurchaseOrderId":4, "PaymentStatus":"CREATED"}`)
	req2, _ := http.NewRequest("POST", url1, bytes.NewBuffer(jsonStr2))

	client2 := &http.Client{}
	resp2, err2 := client2.Do(req2)
	if err2 != nil {
		panic(err2)
	}
	defer resp2.Body.Close()

	if resp2.Status != "200 OK" {
		t.Error("Could not create invoice with this id")
		return
	}

}

//CC10
func TestSubmitInvoiceNonExistPO(t *testing.T) {
	url := "http://localhost:8081/api/invoices"

	var jsonStr = []byte(`{"PurchaseOrderId":4, "PaymentStatus":"CREATED"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status != "400 Bad Request" {
		t.Error("Could not create invoice with this id")
		return
	}

}

//CC11-CC12
func TestApproveInvoice(t *testing.T) {
	url1 := "http://localhost:8081/api/invoices/1/approve"
	url2 := "http://localhost:8081/api/invoices/1"

	var jsonStr = []byte(`{}`)
	req, _ := http.NewRequest("POST", url1, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println(req)
	resp1, err1 := http.Get(url2)
	if err1 != nil {
		t.Error("Problem getting single invoice via REST.")
		return
	}

	pinvoiceJSON, _ := ioutil.ReadAll(resp1.Body)
	var invoice *domain.Invoice
	json.Unmarshal(pinvoiceJSON, &invoice)
	if invoice.PaymentStatus != "PAID" {
		t.Error("Couldn't match with updated PlantReturnDate")
		return
	}

	if resp1.Status != "200 OK" {
		t.Error("Could not get plant hire with this id")
		return
	}

}
