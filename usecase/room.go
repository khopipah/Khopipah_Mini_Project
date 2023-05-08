package usecase

import (
	"errors"
	"fmt"
	"khopipah_mini_project-1/lib/database"
	"khopipah_mini_project-1/models"
)

func CreateRoom(room *models.Room) error {
	// check name cannot be empty
	if room.Name == "" {
		return errors.New("room name cannot be empty")
	}

	//check Class
	if room.Class == "" {
		return errors.New("room class cannot be empty")
	}

	err := database.CreateRoom(room)
	if err != nil {
		return err
	}
	return nil
}

func GetRoom(id uint) (room models.Room, err error) {
	room, err = database.GetRoom(id)
	if err != nil {
		fmt.Println("Error getting room from database")
		return
	}
	return
}

func GetListRooms() (rooms []models.Room, err error) {
	if err != nil {
		fmt.Println("GetListRoom : Error getting user from database")
		return
	}
	return
}

func UpdateRoom(room *models.Room) (err error) {
	err = database.UpdateRoom(room)
	if err != nil {
		fmt.Println("UpdateRoom : Error Updating Room")
		return
	}

	return
}

func DeleteRoom(id uint) (err error) {
	room := models.Room{}
	room.ID = id
	err = database.DeleteRoom(&room)
	if err != nil {
		fmt.Println("DeleteRoom : error deleting room, err ", err)
		return
	}

	return
}
