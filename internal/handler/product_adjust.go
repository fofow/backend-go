package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/msstoci/popow-api/internal/model"
)

func (h *Handler) ProductAdjust(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var res model.GetProductsResponse

	req := model.Product{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	err = h.svc.ProductAdjust(ctx, req)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "success"

	return c.JSON(http.StatusOK, res)
}
