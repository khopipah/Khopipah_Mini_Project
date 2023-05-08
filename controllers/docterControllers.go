package controllers

import (
	"khopipah_mini_project-1/lib/database"
	"khopipah_mini_project-1/models"
	"khopipah_mini_project-1/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetDocterControllers(c echo.Context) error {
	docters, e := database.GetDocters()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all docters",
		"docters": docters,
	})
}

func GetDocterController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	docter, err := database.GetDocter(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes",
		"docters": docter,
	})

}

func CreateDocterController(c echo.Context) error {
	docter := models.Docter{}
	c.Bind(&docter)
	if err := usecase.CreateDocter(&docter); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create docter",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success Create New docters",
		"docters": docter,
	})
}

func DeleteDocterController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := usecase.DeleteDocter(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error elete docter",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf docter tidak dapat dihapus",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "succes Delete docters",
	})
}

func UpdateDocterController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	docter, err := database.GetDocter(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&docter)

	if err := usecase.UpdateDocter(&docter); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "Error Create docter",
			"errorDescription": err,
			"errorMessage":     "Mohon maaf docter tidak dapat diubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update docter",
		"docter":  docter,
	})
}
