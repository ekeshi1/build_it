package test

import (
	"bytes"
	"cs-ut-ee/build-it-project/pkg/internald/domain"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	log "github.com/sirupsen/logrus"
)

var (
	HTTP_URL        = "http://localhost:8081/api/plant-hires"
	HTTP_URL_GET_PH = "http://localhost:8081/api/plant-hires/"
)

func TestUniqueFeatureAutoApproval(t *testing.T) {

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
	req, _ := http.NewRequest("POST", HTTP_URL, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	plantHireJSON, _ := ioutil.ReadAll(resp.Body)
	var plantHire *domain.PlantHire
	json.Unmarshal(plantHireJSON, &plantHire)
	log.Info(plantHire)
	if plantHire.Id == 0 {
		t.Error("Problem creating plant hire.")
		return
	}

	if plantHire.Status != domain.PHApproved {
		t.Error("AUTOAPPROVAL TRIGGERED ON INEXPENSIVE PLANT HIRES")
		return
	}

}

func TestUniqueFeatureNonAutoApproval(t *testing.T) {

	a := `{
		"plantId": 22602061,
		"constructionSiteId": -11449460,
		"supplierId": 82286927,
		"siteEngineerId": -34330736,
		"plantArrivalDate": "2000-02-15T15:04:05Z",
		"plantReturnDate": "2000-08-02T15:04:05Z",
		"plantDailyPrice": 1.0
	}`
	var jsonStr = []byte(a)
	req, _ := http.NewRequest("POST", HTTP_URL, bytes.NewBuffer(jsonStr))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	plantHireJSON, _ := ioutil.ReadAll(resp.Body)
	var plantHire *domain.PlantHire
	json.Unmarshal(plantHireJSON, &plantHire)

	if plantHire.Id == 0 {
		t.Error("Problem creating plant hire.")
		return
	}

	if plantHire.Status != domain.PHCreated {
		t.Error("Autoapproval not triggered on expensive plant hires!")
		return
	}

}
