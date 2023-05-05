package usecase

import (
	"errors"
	"fmt"
	"khopipah_mini_project/lib/database"
	"khopipah_mini_project/models"
)

func CreateUser(user *models.User) error {
	// check name cannot be empty
	if user.Name == "" {
		return errors.New("user name cannot be empty")
	}

	//check email
	if user.Email == "" {
		return errors.New("user Email cannot be empty")
	}

	err := database.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(id uint) (user models.User, err error) {
	user, err = database.GetUser(id)
	if err != nil {
		fmt.Println("Error getting user from database")
		return
	}
	return
}

func GetListUsers() (users []models.User, err error) {
	if err != nil {
		fmt.Println("GetListUser : Error getting user from database")
		return
	}
	return
}

func UpdateUser(user *models.User) (err error) {
	err = database.UpdateUser(user)
	if err != nil {
		fmt.Println("UpdateUser : Error Updating User")
		return
	}

	return
}

func DeleteUser(id uint) (err error) {
	user := models.User{}
	user.ID = id
	err = database.DeleteUser(&user)
	if err != nil {
		fmt.Println("DeleteUser : error deleting user, err ", err)
		return
	}

	return
}
