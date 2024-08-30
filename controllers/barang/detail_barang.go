package barang

import (
	"dkiapp/configs"
	"dkiapp/models/base"
	barangmodel "dkiapp/models/barang"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetBarangDetail(c echo.Context) error {
	// Ambil ID dari parameter URL
	id := c.Param("id")

	var barang barangmodel.Barang
	// Query database untuk mendapatkan barang berdasarkan ID
	result := configs.DB.First(&barang, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusNotFound, base.BaseResponse{
				Status:  false,
				Message: "Barang not found",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed to get data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    barang,
	})
}
