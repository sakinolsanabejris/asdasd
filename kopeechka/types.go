package kopeechka

import "net/http"

type EmailClient struct {
	Client    *http.Client
	ClientKey string
	Domain    string
	Site      string
	OrderId   int
	Status    string `json:"code"`
	Data      struct {
		EmailId string `json:"orderId"` // orderId'nin tipi string, çünkü büyük sayılar olabilir
		Email   string `json:"email"`
		Links   string `json:"links"`
	} `json:"data"`
	Balance float32 `json:"balance"`
	Value   string  `json:"value"`
}
