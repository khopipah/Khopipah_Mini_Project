package database

import (
	"khopipah_mini_project-1/config"
	"khopipah_mini_project-1/models"
)

func CreateDocter(docter *models.Docter) error {
	if err := config.DB.Save(docter).Error; err != nil {
		return err
	}
	return nil
}

func GetDocters() (interface{}, error) {
	var docters []models.Docter

	if err := config.DB.Model(&models.Docter{}).Preload("Pasiens").Find(&docters).Error; err != nil {
		return nil, err
	}
	return docters, nil
}

func GetDocter(id uint) (docter models.Docter, err error) {
	docter.ID = id
	if err = config.DB.Model(&models.Docter{}).Preload("Pasien").First(&docter).Error; err != nil {
		return
	}
	return
}

func UpdateDocter(docter *models.Docter) error {
	if err := config.DB.Updates(docter).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDocter(docter *models.Docter) error {
	if err := config.DB.Delete(docter).Error; err != nil {
		return err
	}
	return nil
}
