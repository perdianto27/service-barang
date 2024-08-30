package barang

import (
	"dkiapp/configs"
	"dkiapp/models/base"
	barangmodel "dkiapp/models/barang"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func DeleteBarangController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var barang barangmodel.Barang
	c.Bind(&barang)

	result := configs.DB.Delete(&barang, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed delete data from database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success deleted",
		Data:    nil,
	})
}
