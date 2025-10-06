package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/fofow/backend-go/internal/model"
)

func (h *Handler) EditProduct(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var res model.Response

	id := c.Param("id")

	intID, _ := strconv.Atoi(id)

	req := model.Product{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	req.ID = int32(intID)

	err = h.svc.EditProduct(ctx, &req)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	return c.JSON(http.StatusOK, res)
}
