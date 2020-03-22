package main

import (
	"ethpaper/ethkey"
	"ethpaper/paperwallet"
	"flag"
	"log"
)

func main() {

	templateFile := flag.String("template", "", "Specify wallet template image")
	outputFile := flag.String("out", "wallet", "Specify paper wallet output path/filename")
	borders := flag.Bool("noborders", false, "Generate QR codes without borders")
	flag.Parse()

	// Generate Ethereum private key
	key := ethkey.NewEthkey()

	// Generate QR codes
	privImg := paperwallet.NewQR(key.Private(), *borders)
	addrImg := paperwallet.NewQR(key.Address(), *borders)

	// Load up template image
	defaultTemplate, err := Asset("assets/paper_wallet_template.png")
	template := paperwallet.NewTemplate(*templateFile, defaultTemplate)

	// Generate paper wallet image from template
	wallet, err := template.Generate(privImg, addrImg, [4]uint8{213, 213, 255, 255}, [4]uint8{213, 255, 246, 255})
	if err != nil {
		log.Fatalln(err)
	}

	paperwallet.SavePng(*outputFile, wallet)
}
