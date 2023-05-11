# QR Code Generator API

Bu API, Go dilinde yazılmıştır ve `go-qrcode` kütüphanesini kullanarak bir URL'yi temsil eden bir QR kodu oluşturur.

## Kurulum

1. Bu depoyu yerel makinenize klonlayın.
2. `go-qrcode` kütüphanesini kurun: `go get github.com/skip2/go-qrcode`
3. `main.go` dosyasını çalıştırın: `go run main.go`

## Kullanım

API, `/generateQR` yoluna bir `GET` isteği yaparak `url` parametresini kullanarak bir QR kodu oluşturur ve `qr.png` dosyasına kaydeder.

Örnek istek: http://localhost:8080/generateQR?url=https://www.example.com

Bu istek, `https://www.example.com` adresini temsil eden bir QR kodu oluşturur ve `qr.png` dosyasına kaydeder.