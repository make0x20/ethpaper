package paperwallet

import (
	"image"
	"image/color"
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

// NewQR returns a new QR image from string
func NewQR(str string, borders bool) image.Image {
	qr, err := qrcode.New(str, qrcode.High)
	if err != nil {
		log.Fatalln(err)
	}

	qr.DisableBorder = borders

	qr.BackgroundColor = color.RGBA{240, 240, 240, 255}
	qr.ForegroundColor = color.RGBA{40, 40, 40, 255}

	return qr.Image(1024)
}
