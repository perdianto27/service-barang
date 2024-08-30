package barang

import (
	"dkiapp/configs"
	"dkiapp/models/base"
	barangmodel "dkiapp/models/barang"
	"net/http"

	"github.com/labstack/echo/v4"
	"time"
)

func strPtr(s string) *string {
	return &s
}

func UpdateBarangController(c echo.Context) error {
	var barangUpdate barangmodel.Barang
	if err := c.Bind(&barangUpdate); err != nil {
		return c.JSON(http.StatusBadRequest, base.BaseResponse{
			Status:  false,
			Message: "Invalid request payload",
			Data:    nil,
		})
	}

	id := c.Param("id")

	var barang barangmodel.Barang
	if result := configs.DB.First(&barang, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, base.BaseResponse{
			Status:  false,
			Message: "Barang not found",
			Data:    nil,
		})
	}

	userID := c.Request().Header.Get("x-user-id")
	barangUpdate.Id = barang.Id
	barangUpdate.CreatedBy = barang.CreatedBy
	updatedAt := time.Now()
	barangUpdate.UpdatedAt = &updatedAt
	barangUpdate.UpdatedBy = strPtr(userID)

	if result := configs.DB.Model(&barang).Updates(barangUpdate); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, base.BaseResponse{
			Status:  false,
			Message: "Failed to update data in database",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, base.BaseResponse{
		Status:  true,
		Message: "Success updated",
		Data:    barangUpdate,
	})
}
