package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetProductByID(c echo.Context) (err error) {
	ctx := c.Request().Context()

	id := c.Param("id")

	intID, _ := strconv.Atoi(id)

	data, err := h.svc.GetProductByID(ctx, int32(intID))
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, data)
	}

	return c.JSON(http.StatusOK, data)
}
