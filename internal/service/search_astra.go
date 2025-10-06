package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/fofow/backend-go/internal/model"
)

func (s *service) SearchAstra(ctx context.Context, input *model.SearchAstraRequest) (res model.SearchAstraResponse, err error) {
	data, err := s.repo.GetAstraBySearch(ctx, input.Search)
	if err != nil {
		return res, err
	}

	// Create the barcode
	qrCode, _ := qr.Encode(data.UUID, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 400, 400)

	var buf bytes.Buffer
	if err := png.Encode(&buf, qrCode); err != nil {
		return res, err
	}
	dataQR := buf.Bytes()

	image := data.Email + ".png"

	err = os.WriteFile(image, dataQR, 0644)
	if err != nil {
		return res, err
	}

	imgBase64Str := base64.StdEncoding.EncodeToString(dataQR)

	res.Data.ID = data.ID
	res.Data.Name = data.Name
	res.Data.Email = data.Email
	res.Data.Phone = data.Phone
	res.Data.Company = data.Company
	res.Data.UUID = data.UUID
	res.Data.QRBase64 = imgBase64Str

	return res, err
}
