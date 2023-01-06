package main

import "github.com/skip2/go-qrcode"

func main() {
	qrcode.WriteFile("http://redseanet.com/", qrcode.Medium, 256, "./golang_qrcode.png")
}
