package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPass(t *testing.T) {
	fmt.Println("Always passing test")
}

func TestGetPlantHire(t *testing.T) {
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
