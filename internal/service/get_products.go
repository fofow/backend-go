package service

import (
	"context"

	"gitlab.com/msstoci/popow-api/internal/model"
)

func (s *service) GetProducts(ctx context.Context) (products []model.Product, err error) {
	products, err = s.repo.GetProducts(ctx)
	if err != nil {
		return products, err
	}

	return products, err
}
