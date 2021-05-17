package repositories

import (
	"cs-ut-ee/build-it-project/pkg/internald/domain"
	"math"

	"fmt"
	"time"

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

func (phr *PlantHireRepository) CalculatePrice(start string, end string, pricePerDay float64) (float64, error) {
	layout := "2006-01-02T15:04:05Z"
	startDate, err := time.Parse(layout, start)

	endDate, err := time.Parse(layout, end)

	if err != nil {
		fmt.Println(err)
	}

	diff := endDate.Sub(startDate)

	if diff < 0 {
		return 0, fmt.Errorf("End date must be bigger than start date")
	}
	diffInDays := math.Ceil(float64(diff/((1e9)*60*60*24))) + 1

	price := diffInDays * pricePerDay
	return price, nil
}
