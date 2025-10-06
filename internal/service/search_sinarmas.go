package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"gitlab.com/msstoci/popow-api/internal/model"
)

func (s *service) SearchSinarmas(ctx context.Context, input *model.SearchSinarmasRequest) (res model.SearchSinarmasResponse, err error) {

	data, err := s.repo.GetSinarmasByPhone(ctx, input.Phone)
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
	res.Data.Telephone = data.Telephone
	res.Data.Attendance = data.Attendance
	res.Data.AttendanceCount = data.AttendanceCount
	res.Data.ProductName = data.ProductName
	res.Data.UUID = data.UUID
	res.Data.QRBase64 = imgBase64Str

	return res, err
}
