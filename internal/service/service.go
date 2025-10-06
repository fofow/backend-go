package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/sirupsen/logrus"
	"github.com/fofow/backend-go/internal/model"
	"github.com/fofow/backend-go/internal/repository"
)

type Service interface {
	GetDataByEmail(ctx context.Context, email string) (res model.Data, err error)
	Register(ctx context.Context, input *model.Data) (err error)
	RegisterSinarmas(ctx context.Context, input *model.RegisterSinarmasRequest) (err error)
	SearchSinarmas(ctx context.Context, input *model.SearchSinarmasRequest) (res model.SearchSinarmasResponse, err error)
	ListDataSearch(ctx context.Context) (res []model.Sinarmas, err error)
	UpdateAttendance(ctx context.Context, id string) (res model.GetDataSinarmasResponse, err error)

	RegisterAstra(ctx context.Context, input *model.RegisterAstraRequest) (err error)
	ListDataAstraSearch(ctx context.Context) (res []model.Astra, err error)
	SearchAstra(ctx context.Context, input *model.SearchAstraRequest) (res model.SearchAstraResponse, err error)
	UpdateAttendanceAstra(ctx context.Context, id string) (res model.GetDataAstraResponse, err error)
	ListIDsAstra(ctx context.Context) (ids []int32, err error)
	UpdateWinnerAstra(ctx context.Context, id int32) (res model.GetDataAstraResponse, err error)

	GetProducts(ctx context.Context) (products []model.Product, err error)
	CreateProducts(ctx context.Context, product *model.Product) (err error)
	EditProduct(ctx context.Context, product *model.Product) (err error)
	GetProductByID(ctx context.Context, id int32) (products model.Product, err error)
	DeleteProductByID(ctx context.Context, id int32) (err error)

	GetProductsActive(ctx context.Context) (products []model.Product, err error)
	ProductAdjust(ctx context.Context, product model.Product) (err error)
}

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &service{repo}
}

func (s *service) GetDataByEmail(ctx context.Context, email string) (res model.Data, err error) {

	res, err = s.repo.GetDataByEmail(ctx, email)
	if err != nil {
		return res, err
	}

	// Create the barcode
	qrCode, _ := qr.Encode("INV00"+fmt.Sprint(res.ID), qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	var buf bytes.Buffer
	if err := png.Encode(&buf, qrCode); err != nil {
		logrus.Error(err)
	}
	data := buf.Bytes()

	imgBase64Str := base64.StdEncoding.EncodeToString(data)

	res.QRBase64 = "data:image/png;base64," + imgBase64Str

	return res, err
}
