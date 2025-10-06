package service

import (
	"context"

	"gitlab.com/msstoci/popow-api/internal/model"
)

func (s *service) UpdateAttendance(ctx context.Context, uuid string) (res model.GetDataSinarmasResponse, err error) {

	data, err := s.repo.GetSinarmasByUUID(ctx, uuid)
	if err != nil {
		return res, err
	}

	tmpData := model.Sinarmas{
		ID:              data.ID,
		Name:            data.Name,
		Email:           data.Email,
		Telephone:       data.Telephone,
		Attendance:      data.Attendance,
		AttendanceCount: data.AttendanceCount,
		ProductName:     data.ProductName,
		RegisterDate:    data.RegisterDate,
		UUID:            uuid,
		QRBase64:        data.QRBase64,
	}

	res.Data = tmpData

	err = s.repo.SetActive(ctx, uuid)

	return res, err
}
