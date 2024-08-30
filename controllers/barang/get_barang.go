package barang

import (
	"dkiapp/configs"
	basemodel "dkiapp/models/base"
	barangmodel "dkiapp/models/barang"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetbBarangController(c echo.Context) error {
	var bs []barangmodel.Barang
	result := configs.DB.Find(&bs)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, basemodel.BaseResponse{
			Status:  false,
			Message: "Failed get data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, basemodel.BaseResponse{
		Status:  true,
		Message: "Success",
		Data:    bs,
	})
}
