package main

import (
	"log"
	"net/http"

	. "github.com/barisakdas/QRCodeGenerator/Handlers"
)

func main() {
	http.HandleFunc("/generateQR", HandleGenerateQR)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
