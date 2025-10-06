package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"html/template"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/msstoci/popow-api/internal/model"
	"gopkg.in/gomail.v2"
)

func (s *service) RegisterAstra(ctx context.Context, input *model.RegisterAstraRequest) (err error) {

	uuid := uuid.New().String()

	input.UUID = uuid
	err = s.repo.InsertAstra(ctx, input)
	if err != nil {
		return err
	}

	// Create the barcode
	qrCode, _ := qr.Encode(uuid, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 400, 400)

	var buf bytes.Buffer
	if err := png.Encode(&buf, qrCode); err != nil {
		return err
	}
	data := buf.Bytes()

	image := input.Email + ".png"

	err = os.WriteFile(image, data, 0644)
	if err != nil {
		return err
	}

	imgBase64Str := base64.StdEncoding.EncodeToString(data)

	registrationNumber := fmt.Sprintf("%0*d", 4, input.ID)

	dataEmail := struct {
		Name               string
		Email              string
		Phone              string
		Company            string
		QRCode             string
		RegistrationNumber string
	}{
		Name:               input.Name,
		Email:              input.Email,
		Phone:              input.Phone,
		Company:            input.Company,
		QRCode:             imgBase64Str,
		RegistrationNumber: registrationNumber,
	}

	// Parse the HTML template
	tmpl, err := template.New("emailTemplate").Parse(htmlBodyRawAstra)
	if err != nil {
		return err
	}

	// Buffer untuk menyimpan hasil dari template yang telah di-render
	var body bytes.Buffer

	// Eksekusi template dengan data
	err = tmpl.Execute(&body, dataEmail)
	if err != nil {
		return err
	}

	go func() {
		errSendMail := sendmailAstra(body.String(), input.Email, image)
		if errSendMail != nil {
			logrus.Error(errSendMail)
		}

		os.Remove("./" + image)
	}()

	return err
}

const htmlBodyRawAstra = `
	<html>
	<head>
	</head>
	<body style="margin: 0; padding: 0; height: 100%;">
		<div id="root" style="background-color: #091B34; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif; font-size: 16px; -webkit-font-smoothing: antialiased; -moz-osx-font-smoothing: grayscale; width: 100%; height: 100%; padding: 25px;">
			<span id="logo" style="display: block; text-align: center; color: #fff; font-size: 30px; font-weight: 600; margin-bottom: 20px;">
				<img src="https://registrasi-astra.vercel.app/new-logo.png" width="500"/>
			</span>
			<div id="content" style="background-color: #fff; border-radius: 5px; padding: 15px; width: 600px; margin: 0 auto;">
				<p>Dear Mr/Mrs. {{.Name}},</p>
				<p>Your reservation for Grand Product Knowledge Agents has been successfully confirmed. Please find your barcode for attendant event :</p>
				<p>Name: {{.Name}}</p>
				<p>Email: {{.Email}}</p>
				<p>No HP: {{.Phone}}</p>
				<p>Company: {{.Company}}</p>
				<br />

				<p>Venue : Altea Blvd Cibubur</p>
				<p>Date : Thursday, September 5th 2024</p>
				<p>Time : 9 AM – 12 AM</p>
				<p>Registration Number : {{.RegistrationNumber}}</p>
				<br />
				<img class="qr" src="cid:{{.Email}}.png" alt="QR Code" style="display: block; width: 324px; margin: 0 auto;">
			</div>
		</div>
	</body>
	</html>
`

const ConfigSmtpHostAstra = "smtp.gmail.com"
const ConfigSmtpPortAstra = 587
const ConfigSenderNameAstra = "Altea BLVD - Astra Land Cibubur <GrandPKalteablvd@gmail.com>"
const ConfigAuthEmailAstra = "GrandPKalteablvd@gmail.com"
const ConfigAuthPasswordAstra = "wypo vzje vfjf sgxh"

func sendmailAstra(htmlBody, email, image string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", ConfigSenderNameAstra)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Reservation Grand Product Knowledge Agents ALTEA BLVD")
	mailer.SetBody("text/html", htmlBody)
	mailer.Embed("./" + image)
	mailer.Attach("./" + image)

	dialer := gomail.NewDialer(
		ConfigSmtpHostAstra,
		ConfigSmtpPortAstra,
		ConfigAuthEmailAstra,
		ConfigAuthPasswordAstra,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}
