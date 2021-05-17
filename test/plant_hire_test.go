package test

import (
	"bytes"
	"cs-ut-ee/build-it-project/pkg/internald/domain"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	HTTP_URL_CREATE_PH = "http://buildit:8081/api/plant-hires"
)

//CC1
func TestCreatePlantHire(t *testing.T) {
	a := `{
		"plantId": 22602061,
		"constructionSiteId": -11449460,
		"supplierId": 82286927,
		"siteEngineerId": -34330736,
		"plantArrivalDate": "2000-02-15T15:04:05Z",
		"plantReturnDate": "2000-03-02T15:04:05Z",
		"plantDailyPrice": 1.0
	}`
	var jsonStr = []byte(a)
	req, _ := http.NewRequest("POST", HTTP_URL_CREATE_PH, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	plantHireJSON, _ := ioutil.ReadAll(resp.Body)
	var plantHire *domain.PlantHire
	json.Unmarshal(plantHireJSON, &plantHire)
	//log.Info(plantHire)
	if plantHire.Id == 0 {
		t.Error("Problem creating plant hire.")
		return
	}

}

//CC2
func TestModifyPlantHireDates(t *testing.T) {
	url := "http://buildit:8081/api/plant-hires/1"

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

//CC4
func TestGetPlantHireById(t *testing.T) {
	resp, err := http.Get("http://buildit:8081/api/plant-hires/1")
	if err != nil {
		t.Error("Problem getting single plant hire via REST.")
		return
	}

	if resp.Status != "200 OK" {
		t.Error("Could not get plant hire with this id")
		return
	}
}

//CC5
func TestModifyPlantHireStatus(t *testing.T) {
	url1 := "http://buildit:8081/api/plant-hires/1/status"
	url2 := "http://buildit:8081/api/plant-hires/1"

	var jsonStr = []byte(`[{"op": "replace", "path": "/status","value": "MODIFIED"}]`)
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
	plantHireJSON, _ := ioutil.ReadAll(resp1.Body)
	var plantHire *domain.PlantHire
	json.Unmarshal(plantHireJSON, &plantHire)

	if plantHire.Status != "MODIFIED" {
		t.Error("Couldn't match with updated PlantArrivalDate")
		return
	}
}

//CC8
func TestModifyPlantHireExtension(t *testing.T) {
	url1 := "http://buildit:8081/api/plant-hires/2/extension"
	url2 := "http://buildit:8081/api/plant-hires/2"

	var jsonStr = []byte(`{"PlantReturnDate":"2021-06-06T00:00:00Z"}`)
	req, _ := http.NewRequest("PUT", url1, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println(resp.Body)
	resp1, err1 := http.Get(url2)
	if err1 != nil {
		t.Error("Problem getting single plant hire via REST.")
		return
	}

	plantHireJSON, _ := ioutil.ReadAll(resp1.Body)
	var plantHire *domain.PlantHire
	json.Unmarshal(plantHireJSON, &plantHire)
	//fmt.Println(plantHire)
	if plantHire.PlantReturnDate != "2021-06-06T00:00:00Z" {
		t.Error("Couldn't match with updated PlantReturnDate")
		return
	}

	if resp1.Status != "200 OK" {
		t.Error("Could not get plant hire with this id")
		return
	}

}
