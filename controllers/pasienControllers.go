package controllers

import (
	"Khopipah_Mini_Project-1/lib/database"
	"Khopipah_Mini_Project-1/models"
	"Khopipah_Mini_Project-1/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetPasienControllers(c echo.Context) error {
	pasien, e := database.GetPasiens()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all pasien",
		"pasien":  pasien,
	})
}

func GetPasienController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	pasien, err := database.GetPasien(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes",
		"pasien":  pasien,
	})

}

func CreatePasienController(c echo.Context) error {
	pasien := models.Pasien{}
	c.Bind(&pasien)
	if err := usecase.CreatePasien(&pasien); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create pasien",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Create New pasiens",
		"pasien":  pasien,
	})
}

func DeletePasienController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := usecase.DeletePasien(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error delete pasien",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf pasien tidak dapat dihapus",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes Delete pasien",
	})
}

func UpdatePasienController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	pasien, err := database.GetPasien(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&pasien)

	if err := usecase.UpdatePasien(&pasien); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create pasien",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf detail pasien tidak dapat diubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update pasien",
		"pasien":  pasien,
	})
}
