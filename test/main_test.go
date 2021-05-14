package test

import (
	"bytes"
	"cs-ut-ee/build-it-project/pkg/internald/domain"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPass(t *testing.T) {
	fmt.Println("Always passing test")
}

func TestModifyPlantHireDates(t *testing.T) {
	url := "http://localhost:8081/api/plant-hires/1"

	var jsonStr = []byte(`[{"op": "replace", "path": "/plantArrivalDate","value": "2022-05-01T00:00:00Z"}]`)
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	resp1, err1 := http.Get(url)
	if err1 != nil {
		t.Error("Problem getting single plant hire via REST.")
		return
	}

	plantHireJSON, _ := ioutil.ReadAll(resp1.Body)
	var plantHire *domain.PlantHire
	json.Unmarshal(plantHireJSON, &plantHire)

	if plantHire.PlantArrivalDate != "2022-05-01T00:00:00Z" {
		t.Error("Couldn't match with updated PlantArrivalDate")
		return
	}
}

func TestGetPlantHireById(t *testing.T) {
	resp, err := http.Get("http://localhost:8081/api/plant-hires/1")
	if err != nil {
		t.Error("Problem getting single plant hire via REST.")
		return
	}

	if resp.Status != "200 OK" {
		t.Error("Could get plant hire with this id")
		return
	}
}

func TestModifyPlantHireStatus(t *testing.T) {
	url1 := "http://localhost:8081/api/plant-hires/1/status"
	url2 := "http://localhost:8081/api/plant-hires/1"

	var jsonStr = []byte(`[{"op": "replace", "path": "/status","value": "MODIFIED"}]`)
	req, _ := http.NewRequest("PATCH", url1, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	defer resp.Body.Close()

	resp1, err1 := http.Get(url2)
	if err1 != nil {
		t.Error("Problem getting single plant hire via REST.")
		return
	}

	plantHireJSON, _ := ioutil.ReadAll(resp1.Body)
	var plantHire *domain.PlantHire
	json.Unmarshal(plantHireJSON, &plantHire)

	if plantHire.Status != "MODIFIED" {
		t.Error("Couldn't match with updated PlantArrivalDate")
		return
	}
}
