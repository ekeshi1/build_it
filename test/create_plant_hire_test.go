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
	HTTP_URL_CREATE_PH = "http://buildit:8080/api/plant-hires"
)

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
	log.Info(plantHire)
	if plantHire.Id == 0 {
		t.Error("Problem creating plant hire.")
		return
	}

}
