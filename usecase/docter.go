package usecase

import (
	"errors"
	"fmt"
	"khopipah_mini_project-1/lib/database"
	"khopipah_mini_project-1/models"
)

func CreateDocter(docter *models.Docter) error {
	// check name cannot be empty
	if docter.Name == "" {
		return errors.New("docter name cannot be empty")
	}

	//check spesialis
	if docter.Spesialis == "" {
		return errors.New("docter spesialis cannot be empty")
	}

	//check gender
	if docter.Gender == "" {
		return errors.New("docter gender cannot be empty")
	}

	//check telp
	if docter.Telp == "" {
		return errors.New("docter telp cannot be empty")
	}

	//check alamat
	if docter.Alamat == "" {
		return errors.New("docter alamat cannot be empty")
	}

	err := database.CreateDocter(docter)
	if err != nil {
		return err
	}
	return nil
}

func GetDocter(id uint) (docter models.Docter, err error) {
	docter, err = database.GetDocter(id)
	if err != nil {
		fmt.Println("Error getting docter from database")
		return
	}
	return
}

func GetListDocters() (docters []models.Docter, err error) {
	if err != nil {
		fmt.Println("GetListDocter : Error getting user from database")
		return
	}
	return
}

func UpdateDocter(docter *models.Docter) (err error) {
	err = database.UpdateDocter(docter)
	if err != nil {
		fmt.Println("UpdateDocter : Error Updating Docter")
		return
	}

	return
}

func DeleteDocter(id uint) (err error) {
	docter := models.Docter{}
	docter.ID = id
	err = database.DeleteDocter(&docter)
	if err != nil {
		fmt.Println("DeleteDocter : error deleting docter, err ", err)
		return
	}

	return
}
