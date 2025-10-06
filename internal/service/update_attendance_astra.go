package service

import (
	"context"

	"gitlab.com/msstoci/popow-api/internal/model"
)

func (s *service) UpdateAttendanceAstra(ctx context.Context, uuid string) (res model.GetDataAstraResponse, err error) {

	data, err := s.repo.GetAstraByUUID(ctx, uuid)
	if err != nil {
		return res, err
	}

	tmpData := model.Astra{
		ID:              data.ID,
		Name:            data.Name,
		Email:           data.Email,
		Phone:       	 data.Phone,
		Company:		 data.Company,
		UUID:            uuid,
		QRBase64:        data.QRBase64,
	}

	res.Data = tmpData

	err = s.repo.SetActiveAstra(ctx, uuid)

	return res, err
}
