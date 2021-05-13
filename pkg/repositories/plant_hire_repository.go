package repositories

import (
	"cs-ut-ee/build-it-project/pkg/internald/domain"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PlantHireRepository struct {
	gormDB *gorm.DB
}

func NewPlantHireRepository(gormDB *gorm.DB) *PlantHireRepository {
	return &PlantHireRepository{
		gormDB: gormDB,
	}
}

func (phr *PlantHireRepository) CreatePlantHire(ph *domain.PlantHire) (*domain.PlantHire, error) {
	var res *gorm.DB
	if res = phr.gormDB.Create(ph); res.Error != nil {
		log.Errorf("Couldn't insert plant hire to db", res.Error)
		return nil, res.Error
	}

	if res.RowsAffected != 1 {
		log.Errorf("Couldn't insert plant hire to db. 0 rows inserted")
		return nil, res.Error
	}

	return ph, nil
}

func (phr *PlantHireRepository) GetPlantHireById(id int64) (*domain.PlantHire, error) {
	var plantHire *domain.PlantHire
	err := phr.gormDB.First(&plantHire, id).Error

	if err != nil {
		log.Errorf("Couldn't get plant hire by its id")
		return nil, err
	}

	return plantHire, nil
}

func (phr *PlantHireRepository) ModifyPlantHire(plantHire *domain.PlantHire, modifiedPlantHire *domain.PlantHire) (*domain.PlantHire, error) {
	err := phr.gormDB.Model(&plantHire).Save(modifiedPlantHire).Error

	if err != nil {
		log.Errorf("Couldn't modify plant hire to db. 0 rows modified")
		return nil, err
	}

	return modifiedPlantHire, nil
}
