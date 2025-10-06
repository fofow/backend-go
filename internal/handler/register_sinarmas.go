package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/msstoci/popow-api/internal/model"
)

func (h *Handler) RegisterSinarmas(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var res model.RegisterSinarmasResponse

	req := model.RegisterSinarmasRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	err = h.svc.RegisterSinarmas(ctx, &req)
	if err != nil {
		res.Data = req
		res.Message = err.Error()
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Data = req

	res.Message = "success"

	return c.JSON(http.StatusOK, res)
}
