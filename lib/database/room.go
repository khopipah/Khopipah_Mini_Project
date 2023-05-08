package database

import (
	"khopipah_mini_project-1/config"
	"khopipah_mini_project-1/models"
)

func CreateRoom(room *models.Room) error {
	if err := config.DB.Save(room).Error; err != nil {
		return err
	}
	return nil
}

func GetRooms() (interface{}, error) {
	var rooms []models.Room

	if err := config.DB.Model(&models.Room{}).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}

func GetRoom(id uint) (room models.Room, err error) {
	room.ID = id
	if err = config.DB.Model(&models.Room{}).Error; err != nil {
		return
	}
	return
}

func UpdateRoom(room *models.Room) error {
	if err := config.DB.Updates(room).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRoom(room *models.Room) error {
	if err := config.DB.Delete(room).Error; err != nil {
		return err
	}
	return nil
}
