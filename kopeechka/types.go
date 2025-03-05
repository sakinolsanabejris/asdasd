package kopeechka

import "net/http"

type EmailClient struct {
	Client    *http.Client
	ClientKey string
	Domain    string
	Site      string
	OrderId   int
	Status    int  `json:"code"`
	EmailId   string  `json:"orderId"`
	Email     string  `json:"email"`
	Balance   float32 `json:"balance"`
	Value     string  `json:"value"`
}
