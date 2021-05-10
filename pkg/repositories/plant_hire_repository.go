package repositories

import (
	"cs-ut-ee/build-it-project/pkg/internald/domain"
	"encoding/json"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch"
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

func (phr *PlantHireRepository) ModifyPlantHire(p []byte, id int64) (*domain.PlantHire, error){
	var plantHire domain.PlantHire
	
	
	phr.gormDB.First(&plantHire, id)
	
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
	
	err = phr.gormDB.Model(&plantHire).Save(modifiedPlantHire).Error

	if err != nil {
		log.Errorf("Couldn't modify plant hire to db. 0 rows modified")
		return nil, err
	}

	return modifiedPlantHire, nil
}
