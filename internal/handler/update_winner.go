package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/fofow/backend-go/internal/model"
)

func (h *Handler) UpdateWinnerAstra(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var res model.GetDataAstraResponse

	id := c.Param("id")

	idInt, _ := strconv.Atoi(id)

	data, err := h.svc.UpdateWinnerAstra(ctx, int32(idInt))
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
