package kopeechka

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (g *EmailClient) BuyEmail(ClientKey string, Domain string, Host string) (string, string) {
	g.ClientKey = ClientKey
	g.Domain = Domain
	g.Site = Host

	client := http.DefaultClient

	// API'ye GET isteği gönder
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.online-disposablemail.com/api/mailbox?apiKey=%v&serviceId=46&emailTypeId=3&quantity=1&linkPriority=true", g.ClientKey), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// API yanıtını çözümle
	var apiResponse ApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Fatal(err)
	}

	// Eğer API'den başarılı yanıt geldiyse
	if apiResponse.Code == 200 && len(apiResponse.Data.Orders) > 0 {
		// Sipariş bilgilerini al
		order := apiResponse.Data.Orders[0]
		g.EmailId = order.OrderId
		g.Email = order.Email
		return g.EmailId, g.Email
	} else {
		return "NOT_OKAY", "ERROR"
	}
}

func (g *EmailClient) GetLetter(ClientKey string, OrderId int) string {
	g.ClientKey = ClientKey
	g.OrderId = OrderId

	client := http.DefaultClient

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.online-disposablemail.com/api/latest/code?orderId=%v", g.OrderId), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Client EmailClient
	err = json.Unmarshal(body, &Client)
	if err != nil {
		log.Fatal(err)
	}

	return Client.Value
}
