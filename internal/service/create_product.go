package service

import (
	"context"

	"github.com/fofow/backend-go/internal/model"
)

func (s *service) CreateProducts(ctx context.Context, product *model.Product) (err error) {

	err = s.repo.InsertProduct(ctx, product)
	if err != nil {
		return err
	}

	return err
}
