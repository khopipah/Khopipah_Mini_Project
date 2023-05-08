package routes

import (
	"khopipah_mini_project-1/controllers"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New() *echo.Echo {
	e := echo.New()
	//route docter
	docter := e.Group("/docters")
	docter.GET("", controllers.GetDocterControllers)
	docter.POST("", controllers.CreateDocterController)
	docter.GET(":id", controllers.GetDocterController)
	docter.PUT(":id", controllers.UpdateDocterController)
	docter.DELETE(":id", controllers.DeleteDocterController)

	//route pasien
	pasien := e.Group("/pasiens")
	pasien.GET("", controllers.GetPasienControllers)
	pasien.POST("", controllers.CreatePasienController)
	pasien.GET(":id", controllers.GetPasienController)
	pasien.PUT(":id", controllers.UpdatePasienController)
	pasien.DELETE(":id", controllers.DeletePasienController)

	//route pasien
	room := e.Group("/rooms")
	room.GET("", controllers.GetRoomControllers)
	room.POST("", controllers.CreateRoomController)
	room.GET(":id", controllers.GetRoomController)
	room.PUT(":id", controllers.UpdateRoomController)
	room.DELETE(":id", controllers.DeleteRoomController)

	return e
}
