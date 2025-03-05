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

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.online-disposablemail.com/api/mailbox?apiKey=913bae74d6224bf6a24dcf354952f1b9&serviceId=46&emailTypeId=3&quantity=1&linkPriority=true"), nil)
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

	var Status EmailClient
	err = json.Unmarshal(body, &Status)
	if err != nil {
		log.Fatal(err)
	}
	if Status.Status == "200" {
		return Status.EmailId, Status.Email
	} else {
		return "NOT_OKAY", "ERROR"
	}
}

func (g *EmailClient) DeleteMail(ClientKey string, OrderId int) (string, error) {
	g.ClientKey = ClientKey
	g.OrderId = OrderId

	client := http.DefaultClient

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.anymessage.shop/email/cancel?token=OKOh6P6sAGYl5BuPel4rAT9hEh6R2VAcn&id=%v", g.OrderId), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == http.StatusOK {
		return "OK", nil
	} else {
		return "NOT_OKAY", nil
	}
}

// func (g *EmailClient) GetEmail(ClientKey string, MailId int) (string, string) {
// 	g.ClientKey = ClientKey
// 	g.OrderId = MailId
// }

func (g *EmailClient) GetBalance(ClientKey string) float32 {
	g.ClientKey = ClientKey

	client := http.DefaultClient

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.anymessage.shop/user/balance?token=%v", g.ClientKey), nil)
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
	if Client.Balance > 0.01 {
		return Client.Balance
	} else {
		return 0
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
