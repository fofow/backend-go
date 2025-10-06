package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/fofow/backend-go/internal/model"
)

func (h *Handler) SearchAstra(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var res model.SearchAstraResponse

	req := model.SearchAstraRequest{}
	search := c.QueryParam("search")
	req.Search = search

	data, err := h.svc.SearchAstra(ctx, &req)
	if err != nil {
		res.Message = err.Error()
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	res = data

	res.Message = "success"

	return c.JSON(http.StatusOK, res)
}
