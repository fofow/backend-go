package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/fofow/backend-go/internal/model"
)

func (h *Handler) CreateProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var res model.Response

	req := model.Product{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	err = h.svc.CreateProducts(ctx, &req)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	return c.JSON(http.StatusOK, res)
}
