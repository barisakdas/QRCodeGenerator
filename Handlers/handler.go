package Handlers

import (
	"net/http"

	"github.com/skip2/go-qrcode"
)

func HandleGenerateQR(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}

	err := qrcode.WriteFile(url, qrcode.Medium, 256, "qr.png")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("QR code generated and saved to qr.png"))
}
