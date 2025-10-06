package service

import (
	"context"

	"github.com/fofow/backend-go/internal/model"
)

func (s *service) ListDataSearch(ctx context.Context) (res []model.Sinarmas, err error) {

	data, err := s.repo.GetDataSinarmas(ctx)
	if err != nil {
		return res, err
	}

	return data, err
}
