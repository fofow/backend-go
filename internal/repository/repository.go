package repository

import (
	"context"

	"gitlab.com/msstoci/popow-api/internal/model"
	"gitlab.com/msstoci/popow-api/pkg/database"
)

type repository struct {
	db *database.Store
}

type Repository interface {
	GetDataByEmail(ctx context.Context, email string) (res model.Data, err error)
	Insert(ctx context.Context, input *model.Data) (err error)
	InsertSinarmas(ctx context.Context, input *model.RegisterSinarmasRequest) (err error)
	InsertAstra(ctx context.Context, input *model.RegisterAstraRequest) (err error)

	GetSinarmasByPhone(ctx context.Context, phone string) (res model.RegisterSinarmasRequest, err error)
	GetDataSinarmas(ctx context.Context) (res []model.Sinarmas, err error)
	SetActive(ctx context.Context, uuid string) (err error)
	GetSinarmasByUUID(ctx context.Context, uuid string) (res model.RegisterSinarmasRequest, err error)
	GetDataAstra(ctx context.Context) (res []model.Astra, err error)
	GetAstraBySearch(ctx context.Context, search string) (res model.Astra, err error)
	SetActiveAstra(ctx context.Context, uuid string) (err error)
	GetAstraByUUID(ctx context.Context, uuid string) (res model.RegisterAstraRequest, err error)
	GetAstraByID(ctx context.Context, id int32) (res model.RegisterAstraRequest, err error)
	GetIDsAstra(ctx context.Context) (ids []int32, err error)
	SetWinnerAstra(ctx context.Context, id int32) (err error)

	GetProducts(ctx context.Context) (products []model.Product, err error)
	InsertProduct(ctx context.Context, input *model.Product) (err error)
	UpdateProduct(ctx context.Context, input *model.Product) (err error)
	GetProductByID(ctx context.Context, id int32) (product model.Product, err error)
	SoftDeleteProductByID(ctx context.Context, id int32) (err error)
}

func New(db *database.Store) Repository {
	return &repository{db}
}

func (r *repository) GetDataByEmail(ctx context.Context, email string) (res model.Data, err error) {

	err = r.db.GetSlave().QueryRowxContext(
		ctx,
		`select id, name, company, phone, email, register_date from user where email = ? order by id desc limit 1`,
		email,
	).StructScan(&res)

	return res, err
}

func (r *repository) Insert(ctx context.Context, input *model.Data) (err error) {
	res, err := r.db.GetMaster().ExecContext(
		ctx,
		`INSERT INTO user (name, company, phone, email) VALUES (?, ?, ?, ?)`,
		input.Name,
		input.Company,
		input.Telephone,
		input.Email,
	)

	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()

	input.ID = int32(id)

	return err
}

func (r *repository) GetProducts(ctx context.Context) (products []model.Product, err error) {
	err = r.db.GetSlave().SelectContext(
		ctx,
		&products,
		`select id, name, quantity from product WHERE deleted_at IS NULL`,
	)

	return products, err
}

func (r *repository) InsertProduct(ctx context.Context, input *model.Product) (err error) {
	res, err := r.db.GetMaster().ExecContext(
		ctx,
		`INSERT INTO product (name, quantity) VALUES (?, ?)`,
		input.Name,
		input.Quantity,
	)

	id, _ := res.LastInsertId()

	input.ID = int32(id)

	return err
}

func (r *repository) GetProductByID(ctx context.Context, id int32) (product model.Product, err error) {
	err = r.db.GetSlave().QueryRowxContext(
		ctx,
		`select id, name, quantity FROM product WHERE id = ? AND deleted_at IS NULL`,
		id,
	).StructScan(&product)

	return product, err
}

func (r *repository) UpdateProduct(ctx context.Context, input *model.Product) (err error) {
	res, err := r.db.GetMaster().ExecContext(
		ctx,
		`UPDATE product SET name = ?, quantity =  ? WHERE id = ?`,
		input.Name,
		input.Quantity,
		input.ID,
	)

	id, _ := res.LastInsertId()

	input.ID = int32(id)

	return err
}

func (r *repository) SoftDeleteProductByID(ctx context.Context, id int32) (err error) {
	_, err = r.db.GetMaster().ExecContext(
		ctx,
		`UPDATE product SET deleted_at = NOW() WHERE id = ?`,
		id,
	)

	return err
}

func (r *repository) InsertSinarmas(ctx context.Context, input *model.RegisterSinarmasRequest) (err error) {
	res, err := r.db.GetMaster().ExecContext(
		ctx,
		`INSERT INTO user (name, phone, email, attendance, attendance_count, product_name, uuid) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		input.Name,
		input.Telephone,
		input.Email,
		input.Attendance,
		input.AttendanceCount,
		input.ProductName,
		input.UUID,
	)

	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()

	input.ID = int32(id)

	return err
}

func (r *repository) GetSinarmasByPhone(ctx context.Context, phone string) (res model.RegisterSinarmasRequest, err error) {

	err = r.db.GetSlave().QueryRowxContext(
		ctx,
		`select id, name, phone, email, attendance, attendance_count, product_name, uuid FROM user WHERE phone = ? OR email = ?`,
		phone,
		phone,
	).StructScan(&res)

	return res, err
}

func (r *repository) SetActive(ctx context.Context, uuid string) (err error) {
	_, err = r.db.GetMaster().ExecContext(
		ctx,
		`UPDATE user SET is_active = 1 WHERE uuid = ?`,
		uuid,
	)

	return err
}

func (r *repository) GetSinarmasByUUID(ctx context.Context, uuid string) (res model.RegisterSinarmasRequest, err error) {
	err = r.db.GetSlave().QueryRowxContext(
		ctx,
		`select id, name, phone, email, attendance, attendance_count, product_name, uuid FROM user WHERE uuid = ?`,
		uuid,
	).StructScan(&res)

	return res, err
}

func (r *repository) InsertAstra(ctx context.Context, input *model.RegisterAstraRequest) (err error) {
	res, err := r.db.GetMaster().ExecContext(
		ctx,
		`INSERT INTO user_astra (name, phone, email, company, uuid) VALUES (?, ?, ?, ?, ?)`,
		input.Name,
		input.Phone,
		input.Email,
		input.Company,
		input.UUID,
	)

	if err != nil {
		return err
	}

	id, _ := res.LastInsertId()

	input.ID = int32(id)

	return err
}

func (r *repository) GetAstraBySearch(ctx context.Context, search string) (res model.Astra, err error) {

	err = r.db.GetSlave().QueryRowxContext(
		ctx,
		`select id, name, phone, email, company, uuid FROM user_astra WHERE phone = ? OR email = ?`,
		search,
		search,
	).StructScan(&res)

	return res, err
}

func (r *repository) SetActiveAstra(ctx context.Context, uuid string) (err error) {
	_, err = r.db.GetMaster().ExecContext(
		ctx,
		`UPDATE user_astra SET is_active = 1 WHERE uuid = ?`,
		uuid,
	)

	return err
}

func (r *repository) GetAstraByUUID(ctx context.Context, uuid string) (res model.RegisterAstraRequest, err error) {
	err = r.db.GetSlave().QueryRowxContext(
		ctx,
		`select id, name, phone, email, company, uuid FROM user_astra WHERE uuid = ?`,
		uuid,
	).StructScan(&res)

	return res, err
}

func (r *repository) GetIDsAstra(ctx context.Context) (ids []int32, err error) {

	err = r.db.GetSlave().SelectContext(
		ctx,
		&ids,
		`SELECT id FROM user_astra WHERE is_active = 1 AND is_winner = 0`,
	)

	return ids, err
}

func (r *repository) SetWinnerAstra(ctx context.Context, id int32) (err error) {
	_, err = r.db.GetMaster().ExecContext(
		ctx,
		`UPDATE user_astra SET is_winner = 1 WHERE id = ?`,
		id,
	)

	return err
}

func (r *repository) GetAstraByID(ctx context.Context, id int32) (res model.RegisterAstraRequest, err error) {
	err = r.db.GetSlave().QueryRowxContext(
		ctx,
		`select id, name, phone, email, company, uuid FROM user_astra WHERE id = ?`,
		id,
	).StructScan(&res)

	return res, err
}
