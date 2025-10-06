package service

import (
	"context"

	"github.com/fofow/backend-go/internal/model"
)

func (s *service) GetProductByID(ctx context.Context, id int32) (products model.Product, err error) {

	products, err = s.repo.GetProductByID(ctx, id)
	if err != nil {
		return products, err
	}

	return products, err
}
