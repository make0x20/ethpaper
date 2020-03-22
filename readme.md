# Ethpaper

Ethpaper is a simple Ethereum paper wallet generator with an option to specify custom templates

**Ethpaper has not been tested extensively and was created as a hobby project, use at your own risk when using with real Ethereum assets!**

## Features

- Option to specify your own 'template.png' image
- Completely offline - *Ideally, paper wallets are to be generated on air-gapped computers without any connection*
- No dependencies, just run it

## Install
```go get github.com/make0x20/ethpaper```

## Usage
Help `ethpaper -h`: 
```
Usage of ./ethpaper:
  -noborders
    	Generate QR codes without borders
  -out string
    	Specify paper wallet output path/filename (default "wallet")
  -template string
    	Specify wallet template image (default "assets/paper_wallet_template.png")
```
