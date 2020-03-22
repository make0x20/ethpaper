package paperwallet

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/nfnt/resize"
)

type WalletTemplate struct {
	Image image.Image
}

// findAreaByColor finds rectangular region by color
// color parameter must be in rgba format - Example: [4]uint8{255,255,255,0}
func (wt WalletTemplate) findAreasByColor(colorNRGBA color.NRGBA) []image.Rectangle {
	min, max := wt.Image.Bounds().Min, wt.Image.Bounds().Max
	areas := []image.Rectangle{}

	colorStartAt := image.Point{}
	colorEndAt := image.Point{}
	// Loop through all pixels
	for col := min.X; col < max.X; col++ {
		rowHasColor := false
		for row := min.Y; row < max.Y; row++ {
			pixel := wt.Image.At(col, row)

			// Skip if pixel is not the same as colorNRGBA
			if pixel != colorNRGBA {
				continue
			}

			// Set to true if column has the color
			rowHasColor = true

			// Set first pixel location if not set
			if (image.Point{}) == colorStartAt {
				colorStartAt.X = col
				colorStartAt.Y = row
			}

			// Add each new pixel with color to colorEndAt
			colorEndAt.X = col
			colorEndAt.Y = row
		}

		// If column has no color and colorStartAt and colorEndAt were already found, create image.Rectangle
		if !rowHasColor && (image.Point{}) != colorStartAt && (image.Point{}) != colorEndAt {
			// Add area to areas
			rect := image.Rectangle{
				colorStartAt,
				colorEndAt,
			}
			areas = append(areas, rect)
			// Reset points
			colorStartAt, colorEndAt = image.Point{}, image.Point{}
		}
	}

	return areas
}

// Generate populates template with QR codes
// privateColor and addressColor parameters must be in rgba format - Example: [4]uint8{255,255,255,0}
func (wt WalletTemplate) Generate(privateQR, addressQR image.Image, privateColor, addressColor [4]uint8) (image.Image, error) {
	privRects := wt.findAreasByColor(color.NRGBA{privateColor[0], privateColor[1], privateColor[2], privateColor[3]})
	addrRects := wt.findAreasByColor(color.NRGBA{addressColor[0], addressColor[1], addressColor[2], addressColor[3]})

	// Create a new image and draw the template in
	wallet := image.NewRGBA(wt.Image.Bounds())
	draw.Draw(wallet, wallet.Bounds(), wt.Image, image.ZP, draw.Src)

	// Draw all private areas with private key QR
	for _, rect := range privRects {
		resizedQR := resize.Resize(uint(rect.Bounds().Dx()), uint(rect.Bounds().Dy()), privateQR, resize.Lanczos3)
		draw.Draw(wallet, rect.Bounds(), resizedQR, image.ZP, draw.Src)
	}

	// Draw all address areas with address key QR
	for _, rect := range addrRects {
		resizedQR := resize.Resize(uint(rect.Bounds().Dx()), uint(rect.Bounds().Dy()), addressQR, resize.Lanczos3)
		draw.Draw(wallet, rect.Bounds(), resizedQR, image.ZP, draw.Over)
	}

	return wallet, nil
}

// SavePng saves image as PNG
func SavePng(name string, img image.Image) {
	file, err := os.Create(name + ".png")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	png.Encode(file, img)
	fmt.Println("Saved:", name+".png")
}

// NewTemplate returns the paper wallet template image
func NewTemplate(path string, defaultTemplate []byte) WalletTemplate {

	templateFile := bytes.NewReader(defaultTemplate)
	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		template, err := ioutil.ReadAll(file)
		templateFile = bytes.NewReader(template)
	}

	img, err := png.Decode(templateFile)
	if err != nil {
		log.Fatalln(err)
	}

	return WalletTemplate{img}
}
