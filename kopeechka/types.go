package kopeechka

import "net/http"

type ApiResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Orders []struct {
			OrderId string `json:"orderId"`
			Email   string `json:"email"`
		} `json:"orders"`
		Links string `json:"links"`
	} `json:"data"`
}

type EmailClient struct {
	Client    *http.Client `json:"-"` // HTTP client (JSON'dan gelmez)
	ClientKey string       `json:"clientKey"`
	Domain    string       `json:"domain"`
	Site      string       `json:"site"`
	Status    int       `json:"status"`  // API yanıtındaki durum (OK, ERROR vb.)
	Balance   float32      `json:"balance"` // Hesap bakiyesi
	Value     string       `json:"value"`   // E-posta değeri
	EmailId   string       `json:"emailId"` // E-posta kimliği
	Email     string       `json:"email"`   // E-posta adresi
	OrderId   int          `json:"orderId"` // Sipariş kimliği
}
