package services

import (
	"cs-ut-ee/build-it-project/pkg/internald/domain"
	"cs-ut-ee/build-it-project/pkg/internald/ports"

	log "github.com/sirupsen/logrus"
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
