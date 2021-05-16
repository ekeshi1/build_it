package test

import (
	"bytes"
	"net/http"
	"testing"
)

//CC6
func TestCreatePO(t *testing.T) {
	url1 := "http://localhost:8081/api/plant-hires/2/status"
	url2 := "http://localhost:8081/api/purchase-orders/2"

	var jsonStr = []byte(`[{"op": "replace", "path": "/status","value": "APPROVED"}]`)
	req, _ := http.NewRequest("PATCH", url1, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	resp1, err1 := http.Get(url2)
	if err1 != nil {
		t.Error("Problem getting single plant hire via REST.")
		return
	}
	if resp1.Status != "200 OK" {
		t.Error("Could not get purchase order by this plant hire id")
		return
	}
}

//CC7
func TestGetAllPurchaseOrders(t *testing.T) {
	url := "http://localhost:8081/api/purchase-orders"

	resp, err := http.Get(url)
	if err != nil {
		t.Error("Problem getting purchase orders hire via REST.")
		return
	}
	if resp.Status != "200 OK" {
		t.Error("Could not get purchase orders ")
		return
	}
}
