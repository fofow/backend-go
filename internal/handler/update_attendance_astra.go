package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/fofow/backend-go/internal/model"
)

func (h *Handler) UpdateAttendanceAstra(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var res model.GetDataAstraResponse

	id := c.Param("id")

	data, err := h.svc.UpdateAttendanceAstra(ctx, id)
	if err != nil {
		res.Data.UUID = id
		res.Message = err.Error()
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Data = data.Data

	res.Message = "success"

	return c.JSON(http.StatusOK, res)
}
