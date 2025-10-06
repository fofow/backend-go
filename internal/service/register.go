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
)

func (s *service) Register(ctx context.Context, input *model.Data) (err error) {

	err = s.repo.Insert(ctx, input)
	if err != nil {
		return err
	}

	// Create the barcode
	qrCode, _ := qr.Encode("INV00"+fmt.Sprint(input.ID), qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	var buf bytes.Buffer
	if err := png.Encode(&buf, qrCode); err != nil {
		logrus.Error(err)
	}
	data := buf.Bytes()

	imgBase64Str := base64.StdEncoding.EncodeToString(data)

	input.QRBase64 = "data:image/png;base64," + imgBase64Str

	return err
}
