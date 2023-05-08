package usecase

import (
	"errors"
	"fmt"
	"khopipah_mini_project-1/lib/database"
	"khopipah_mini_project-1/models"
)

func CreatePasien(pasien *models.Pasien) error {
	// check name cannot be empty
	if pasien.Name == "" {
		return errors.New("pasien name cannot be empty")
	}

	//check penyakit
	if pasien.Penyakit == "" {
		return errors.New("pasien penyakit cannot be empty")
	}

	err := database.CreatePasien(pasien)
	if err != nil {
		return err
	}
	return nil
}

func GetPasien(id uint) (pasien models.Pasien, err error) {
	pasien, err = database.GetPasien(id)
	if err != nil {
		fmt.Println("Error getting pasien from database")
		return
	}
	return
}

func GetListPasien() (pasien []models.Pasien, err error) {
	if err != nil {
		fmt.Println("GetListPasien : Error getting user from database")
		return
	}
	return
}

func UpdatePasien(pasien *models.Pasien) (err error) {
	err = database.UpdatePasien(pasien)
	if err != nil {
		fmt.Println("UpdatePasien : Error Updating pasien")
		return
	}

	return
}

func DeletePasien(id uint) (err error) {
	Pasien := models.Pasien{}
	Pasien.ID = id
	err = database.DeletePasien(&Pasien)
	if err != nil {
		fmt.Println("DeletePasien : error deleting pasien, err ", err)
		return
	}

	return
}
