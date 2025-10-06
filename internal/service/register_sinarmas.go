package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gitlab.com/msstoci/popow-api/internal/model"
	"gopkg.in/gomail.v2"
)

const htmlBodyRaw = `
	<html>
	<head>
	</head>
	<body style="margin: 0; padding: 0; height: 100%;">
		<div id="root" style="background-color: #4663ac; font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif; font-size: 16px; -webkit-font-smoothing: antialiased; -moz-osx-font-smoothing: grayscale; width: 100%; height: 100%; padding: 25px;">
			<span id="logo" style="display: block; text-align: center; color: #fff; font-size: 30px; font-weight: 600; margin-bottom: 20px;">
				<img src="https://registrasi-sinarmas.vercel.app/new-logo.png" width="500"/>
			</span>
			<div id="content" style="background-color: #fff; border-radius: 5px; padding: 15px; width: 600px; margin: 0 auto;">
				<p>Kepada Tuan/Nyonya. {{.Name}},</p>
				<p>Reservasi Anda untuk acara The Ultimate Phase of Tresor telah berhasil dikonfirmasi. Silakan tunjukan kode barcode Anda untuk menghadiri acara:</p>
				<p>Name: {{.Name}}</p>
				<p>Kantor Agent: {{.ProductName}}</p>
				<p>No Telephone: {{.Phone}}</p>
				<p>Email: {{.Email}}</p>

				<br />

				<p>Tempat : Green Office Park 9 Theater BSD City</p>
				<p>Tanggal : Senin, 21 Juli 2025</p>
				<p>Waktu : 14.00-16.00 WIB</p>
				<p>Nomor Registrasi : {{.RegistrationNumber}}</p>
				<br />
				<img class="qr" src="cid:{{.Email}}.png" alt="QR Code" style="display: block; width: 324px; margin: 0 auto;">
				<p>Catatan:</p>
				<ul> 
					<li>Registrasi dibuka pukul 13.00 WIB</li>
					<li>⁠Undangan berlaku untuk 1 orang</li>
				</ul>
				</div>
		</div>
	</body>
	</html>
`

func (s *service) RegisterSinarmas(ctx context.Context, input *model.RegisterSinarmasRequest) (err error) {

	uuid := uuid.New().String()

	input.UUID = uuid
	err = s.repo.InsertSinarmas(ctx, input)
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

	image := uuid + ".png"

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
		Attendance         string
		AttendanceCount    int32
		ProductName        string
		QRCode             string
		RegistrationNumber string
	}{
		Name:               input.Name,
		Email:              input.Email,
		Phone:              input.Telephone,
		Attendance:         input.Attendance,
		AttendanceCount:    input.AttendanceCount,
		ProductName:        input.ProductName,
		QRCode:             imgBase64Str,
		RegistrationNumber: registrationNumber,
	}

	// Parse the HTML template
	tmpl, err := template.New("emailTemplate").Parse(htmlBodyRaw)
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

		message := `Kepada Tuan/Nyonya. ` + input.Name + `,
Reservasi Anda untuk acara The Ultimate Phase of Tresor telah berhasil dikonfirmasi. Silakan tunjukan kode barcode Anda untuk menghadiri acara:

Nama : ` + input.Name + `
Kantor Agent : ` + input.ProductName + `
No HP : ` + input.Telephone + `
Email : ` + input.Email + `
Tempat : Green Office Park 9 Theater BSD City
Tanggal : Senin, 21 Juli 2025
Waktu : 14.00-16.00 WIB
Nomor Registrasi : ` + registrationNumber + `

Catatan :
•⁠  ⁠Registrasi dibuka pukul 13.00 WIB
•⁠  ⁠Undangan berlaku untuk 1 orang
`

		errSendWa := sendWAImage(input.Telephone, message, image)
		if errSendWa != nil {
			logrus.Error(errSendWa)
		}

		errSendMail := sendmail(body.String(), input.Email, image)
		if errSendMail != nil {
			logrus.Error(errSendMail)
		}

		os.Remove("./" + image)
	}()

	return err
}

// const ConfigSmtpHost = "smtp.office365.com"
// const ConfigSmtpPort = 587
// const ConfigSenderName = "Sinarmas Land <marcomm@sinarmasland.com>"
// const ConfigAuthEmail = "marcomm@sinarmasland.com"
// const ConfigAuthPassword = "B$d2024@"

const ConfigSmtpHost = "smtp.gmail.com"
const ConfigSmtpPort = 587
const ConfigSenderName = "Sinarmas Land <sinarmasland.infinite2025@gmail.com>"
const ConfigAuthEmail = "sinarmasland.infinite2025@gmail.com"
const ConfigAuthPassword = "gzxs kcsn njle dtkh"

func sendmail(htmlBody, email, image string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", ConfigSenderName)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Reservation Media Launch & Prestigious Introduction")
	mailer.SetBody("text/html", htmlBody)
	mailer.Embed("./" + image)
	mailer.Attach("./" + image)

	dialer := gomail.NewDialer(
		ConfigSmtpHost,
		ConfigSmtpPort,
		ConfigAuthEmail,
		ConfigAuthPassword,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}

type FonnteResponse struct {
	Reason string `json:"reason"`
	Status bool   `json:"status"`
}

type FonnteRequest struct {
	Target  string `json:"target"`
	Message string `json:"message"`
}

func sendWa(phoneNumber, message string) error {
	url := "https://api.fonnte.com/send"

	payload, err := json.Marshal(FonnteRequest{
		Target:  phoneNumber,
		Message: message,
	})

	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "9Acn4VP8grschn4yzzep")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var resp FonnteResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("error " + http.StatusText(res.StatusCode))
	}

	if !resp.Status {
		return errors.New(resp.Reason)
	}

	return nil
}

func sendWAImage(phoneNumber, message, path string) (err error) {
	url := "https://api.fonnte.com/send"

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add the fields
	writer.WriteField("target", phoneNumber)
	writer.WriteField("message", message)

	// Add the file
	file, err := os.Open(path) // replace with your actual file path
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", path)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	// Close the writer to finalize the form data
	err = writer.Close()
	if err != nil {
		return err
	}

	// Prepare the HTTP request
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		return err
	}

	// Set the headers
	req.Header.Set("Authorization", "9Acn4VP8grschn4yzzep") // replace with your actual token
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	logrus.Info(string(body))

	return nil
}
