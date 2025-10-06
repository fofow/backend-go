package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/fofow/backend-go/internal/model"
)

func (h *Handler) DeleteProductByID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var res model.Response

	id := c.Param("id")

	intID, _ := strconv.Atoi(id)

	err = h.svc.DeleteProductByID(ctx, int32(intID))
	if err != nil {
		res.Message = err.Error()
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "success"

	return c.JSON(http.StatusOK, res)
}
