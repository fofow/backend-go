package service

import (
	"context"

	"github.com/fofow/backend-go/internal/model"
)

func (s *service) EditProduct(ctx context.Context, product *model.Product) (err error) {

	err = s.repo.UpdateProduct(ctx, product)
	if err != nil {
		return err
	}

	return err
}
