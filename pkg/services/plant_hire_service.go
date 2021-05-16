package services

import (
	"cs-ut-ee/build-it-project/pkg/internald/domain"
	"cs-ut-ee/build-it-project/pkg/internald/ports"

	log "github.com/sirupsen/logrus"

	"encoding/json"
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
)

type PlantHireService struct {
	plantHireRepository ports.PlantHireRepositoryPort
	poDriverPort        ports.PurchaseOrderServiceDriverPort
}

func NewPlantHireService(phr ports.PlantHireRepositoryPort, poDriver ports.PurchaseOrderServiceDriverPort) *PlantHireService {
	return &PlantHireService{
		plantHireRepository: phr,
		poDriverPort:        poDriver,
	}
}

func (s *PlantHireService) CreatePlantHire(ph *domain.PlantHire) (*domain.PlantHire, error) {
	totalPrice, err := s.plantHireRepository.CalculatePrice(ph.PlantArrivalDate, ph.PlantReturnDate, ph.PlantDailyPrice)
	if err != nil {
		log.Errorf("Couldn't create new plant with error: ", err)
		return nil, err
	}
	log.Info(totalPrice)
	ph.PlantTotalPrice = totalPrice
	if totalPrice < 100 {
		ph.Status = domain.PHApproved
	} else {
		ph.Status = domain.PHCreated
	}

	createdPlantHire, err1 := s.plantHireRepository.CreatePlantHire(ph)

	if err1 != nil {
		log.Errorf("Couldn't create new plant with error: ", err1)
		return nil, err1
	}

	log.Debugf("Created plant with id : ", createdPlantHire.Id)
	return createdPlantHire, nil
}

func (s *PlantHireService) ModifyPlantHire(p []byte, id int64) (*domain.PlantHire, error) {

	plantHire, err := s.plantHireRepository.GetPlantHireById(id)

	if err != nil {
		log.Errorf("Couldn't get plant with error: ", err)
		return nil, err
	}

	plantBytes, err := json.Marshal(plantHire)
	if err != nil {
		fmt.Println("Error creating patch json ", err.Error())
		return nil, err
	}
	fmt.Println(string(plantBytes))

	patch, err := jsonpatch.DecodePatch(p)
	if err != nil {
		fmt.Println("Error Decoding patch json ", err.Error())
		return nil, err
	}

	modified, err := patch.Apply(plantBytes)
	if err != nil {
		fmt.Println("Error applying patch json ", err.Error())
		return nil, err
	}
	var modifiedPlantHire *domain.PlantHire
	json.Unmarshal(modified, &modifiedPlantHire)

	mph, err1 := s.plantHireRepository.ModifyPlantHire(plantHire, modifiedPlantHire)
	if err1 != nil {
		log.Errorf("Couldn't update plant with error: ", err)
		return nil, err
	}
	log.Debugf("Created modified plant with id : ", modifiedPlantHire.Id)
	return mph, nil
}

func (s *PlantHireService) GetPlantHireById(id int64) (*domain.PlantHire, error) {
	plantHire, err := s.plantHireRepository.GetPlantHireById(id)

	if err != nil {
		log.Errorf("Couldn't get plant with error: ", err)
		return nil, err
	}
	return plantHire, nil
}

func (s *PlantHireService) ModifyPlantHireExtension(id int64, p *domain.PlantHireExtensionDTO) (*domain.PlantHire, error) {

	plantHire, err := s.plantHireRepository.GetPlantHireById(id)

	if err != nil {
		log.Errorf("Couldn't get plant with error: ", err)
		return nil, err
	}

	var modifiedPlantHire *domain.PlantHire
	modifiedPlantHire = plantHire
	modifiedPlantHire.PlantReturnDate = p.PlantReturnDate
	fmt.Println(plantHire.PlantArrivalDate)
	fmt.Println(p.PlantReturnDate)
	totalPrice, err := s.plantHireRepository.CalculatePrice(plantHire.PlantArrivalDate, p.PlantReturnDate, 127.00)
	plantHire.PlantTotalPrice = totalPrice
	fmt.Println(totalPrice)
	fmt.Println(plantHire)
	fmt.Println(modifiedPlantHire)
	if err != nil {
		log.Errorf("Couldn't create new plant with error: ", err)
		return nil, err
	}

	mph, err1 := s.plantHireRepository.ModifyPlantHire(plantHire, modifiedPlantHire)
	if err1 != nil {
		log.Errorf("Couldn't update plant with error: ", err)
		return nil, err
	}
	log.Debugf("Modify plant hire with id : ", modifiedPlantHire.Id)

	log.Debugf("Trying to send modified po to supplier( RENT IT)")

	var modifiedPurchaseOrder *domain.PurchaseOrder
	//make http request to rent it
	//this can be changed for another transport layer, only by implementing  PurchaseOrderServiceDriverPort interface.
	if isSuccessfull, err := s.poDriverPort.ModifyPurchaseOrder(modifiedPurchaseOrder); err != nil || isSuccessfull != true {
		log.Error("Somethihng went wrong notifying third party about po. Error %v", err)
		return nil, err
	}

	log.Debugf("Modified plant with id : ", modifiedPlantHire.Id)
	return mph, nil
}
