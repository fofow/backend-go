package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()

	data, err := h.svc.GetProducts(ctx)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, data)
	}

	return c.JSON(http.StatusOK, data)
}
