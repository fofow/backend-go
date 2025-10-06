package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/msstoci/popow-api/internal/model"
)

func (h *Handler) ListDataAstra(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var res model.GetDatasAstraResponse

	data, err := h.svc.ListDataAstraSearch(ctx)
	if err != nil {
		res.Message = err.Error()
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Message = "success"

	return c.JSON(http.StatusOK, data)
}
