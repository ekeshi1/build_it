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
}

func NewPlantHireService(phr ports.PlantHireRepositoryPort) *PlantHireService {
	return &PlantHireService{
		plantHireRepository: phr,
	}
}

func (s *PlantHireService) CreatePlantHire(ph *domain.PlantHire) (*domain.PlantHire, error) {

	createdPlantHire, err := s.plantHireRepository.CreatePlantHire(ph)

	if err != nil {
		log.Errorf("Couldn't create new plant with error: ", err)
		return nil, err
	}

	log.Debugf("Created plant with id : ", createdPlantHire.Id)
	return createdPlantHire, nil
}

func(s *PlantHireService) ModifyPlantHire(p []byte, id int64) (*domain.PlantHire, error){
	
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