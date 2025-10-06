package service

import (
	"context"

	"gitlab.com/msstoci/popow-api/internal/model"
)

func (s *service) UpdateWinnerAstra(ctx context.Context, id int32) (res model.GetDataAstraResponse, err error) {

	data, err := s.repo.GetAstraByID(ctx, id)
	if err != nil {
		return res, err
	}

	tmpData := model.Astra{
		ID:       id,
		Name:     data.Name,
		Email:    data.Email,
		Phone:    data.Phone,
		Company:  data.Company,
		UUID:     data.UUID,
		QRBase64: data.QRBase64,
	}

	res.Data = tmpData

	err = s.repo.SetWinnerAstra(ctx, id)

	return res, err
}
