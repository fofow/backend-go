package service

import (
	"context"

	"github.com/fofow/backend-go/internal/model"
)

func (s *service) ListDataAstraSearch(ctx context.Context) (res []model.Astra, err error) {

	data, err := s.repo.GetDataAstra(ctx)
	if err != nil {
		return res, err
	}

	return data, err
}
