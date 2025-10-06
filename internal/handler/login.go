package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"gitlab.com/msstoci/popow-api/internal/model"
)

func (h *Handler) GetDataByEmail(c echo.Context) (err error) {
	ctx := c.Request().Context()

	var res model.Response

	req := model.Login{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	data, err := h.svc.GetDataByEmail(ctx, req.Email)
	if err != nil {
		if data.Name == "" && err == sql.ErrNoRows {
			res.Message = "data is not found"
		}

		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, res)
	}

	res.Data = data
	res.Message = "success"

	return c.JSON(http.StatusOK, res)
}
