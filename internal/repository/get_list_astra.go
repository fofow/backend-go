package repository

import (
	"context"

	"github.com/fofow/backend-go/internal/model"
)

func (r *repository) GetDataAstra(ctx context.Context) (res []model.Astra, err error) {
	err = r.db.GetSlave().SelectContext(
		ctx,
		&res,
		`
			SELECT 
				id, 
				name, 
				phone, 
				email, 
				company,
				uuid 
			FROM 
				user_astra
		`,
	)

	return res, err
}
