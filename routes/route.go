package routes

import (
	"prakerja/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {

	e := echo.New()
	e.GET("/students", controllers.GetAllSiswa)
	e.GET("/students/:id", controllers.GetSiswaByID)
	e.POST("/students", controllers.CreateSiswaController)
	e.DELETE("/students/:id", controllers.DeleteSiswaController)
	e.PUT("/students/:id", controllers.UpdateSiswaController)
	return e

}
