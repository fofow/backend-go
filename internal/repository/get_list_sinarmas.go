package repository

import (
	"context"

	"github.com/fofow/backend-go/internal/model"
)

func (r *repository) GetDataSinarmas(ctx context.Context) (res []model.Sinarmas, err error) {
	err = r.db.GetSlave().SelectContext(
		ctx,
		&res,
		`
			SELECT 
				id, 
				name, 
				phone, 
				email, 
				attendance, 
				attendance_count, 
				product_name, 
				uuid 
			FROM 
				user
		`,
	)

	return res, err
}
