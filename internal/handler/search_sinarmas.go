package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/fofow/backend-go/internal/model"
)

func (h *Handler) SearchSinarmas(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var res model.SearchSinarmasResponse

	req := model.SearchSinarmasRequest{}
	phone := c.QueryParam("phone")
	req.Phone = phone

	data, err := h.svc.SearchSinarmas(ctx, &req)
	if err != nil {
		res.Message = err.Error()
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	res = data

	res.Message = "success"

	return c.JSON(http.StatusOK, res)
}
