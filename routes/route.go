package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	barangcontroller "dkiapp/controllers/barang"
)

func InitRoutes(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.POST("/barang", barangcontroller.AddBarangController)
	e.GET("/list-barang", barangcontroller.GetbBarangController)
	e.GET("/barang/:id", barangcontroller.GetBarangDetail)
	e.PUT("/barang/:id", barangcontroller.UpdateBarangController)
	e.DELETE("/barang/:id", barangcontroller.DeleteBarangController)
}
