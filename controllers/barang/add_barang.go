package barang

import (
	"dkiapp/configs"
	"dkiapp/models/base"
	barangmodel "dkiapp/models/barang"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddBarangController(c echo.Context) error {
	var barang barangmodel.Barang
	c.Bind(&barang)
	userID := c.Request().Header.Get("x-user-id")
	barang.CreatedBy = userID

	barang.UpdatedAt = nil
	barang.UpdatedBy = nil

	result := configs.DB.Create(&barang)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed add data to database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success created",
		Data:    barang,
	})
}
