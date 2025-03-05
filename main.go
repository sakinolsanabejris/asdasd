package main

import (
	"fmt"

	"github.com/sakinolsanabejris/asdasd/kopeechka"
)

func main() {
	// Email Client!
	var emailkey string = "" // Users Key

	client := new(kopeechka.EmailClient)

	// Buys Email Taking 3 arguements, returning mailid and email
	MailId, Email := client.BuyEmail(emailkey, "mail.ru", "twitch.tv")
	fmt.Printf("%v | %v\n", MailId, Email)

	// Cancels Mail

	// Get Letter
	Value := client.GetLetter(emailkey, 1441288813)
	fmt.Printf("%v\n", Value)
}
