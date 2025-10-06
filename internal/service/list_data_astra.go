package service

import (
	"context"

	"gitlab.com/msstoci/popow-api/internal/model"
)

func (s *service) ListDataAstraSearch(ctx context.Context) (res []model.Astra, err error) {

	data, err := s.repo.GetDataAstra(ctx)
	if err != nil {
		return res, err
	}

	return data, err
}
