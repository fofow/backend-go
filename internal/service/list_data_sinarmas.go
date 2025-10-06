package service

import (
	"context"

	"gitlab.com/msstoci/popow-api/internal/model"
)

func (s *service) ListDataSearch(ctx context.Context) (res []model.Sinarmas, err error) {

	data, err := s.repo.GetDataSinarmas(ctx)
	if err != nil {
		return res, err
	}

	return data, err
}
