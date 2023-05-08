package database

import (
	"Khopipah_Mini_Project-1/config"
	"Khopipah_Mini_Project-1/models"
)

func CreatePasien(pasien *models.Pasien) error {
	if err := config.DB.Save(pasien).Error; err != nil {
		return err
	}
	return nil
}

func GetPasiens() (interface{}, error) {
	var pasiens []models.Pasien

	if err := config.DB.Model(&models.Pasien{}).Error; err != nil {
		return nil, err
	}
	return pasiens, nil
}

func GetPasien(id uint) (pasien models.Pasien, err error) {
	pasien.ID = id
	if err = config.DB.Model(&models.Pasien{}).Error; err != nil {
		return
	}
	return
}

func UpdatePasien(pasien *models.Pasien) error {
	if err := config.DB.Updates(pasien).Error; err != nil {
		return err
	}
	return nil
}

func DeletePasien(pasien *models.Pasien) error {
	if err := config.DB.Delete(pasien).Error; err != nil {
		return err
	}
	return nil
}
