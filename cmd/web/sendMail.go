package main

import (
	"github.com/AnonymFromInternet/Motel/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"time"
)

func listenForMail() {
	// Asynchronous goroutine
	go func() {
		for {
			message := <-appConfig.MailChan
			// Will called only if message gets the value???
			sendMessage(message)
		}
	}()
}

func sendMessage(mailData models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 15 * time.Second
	server.SendTimeout = 15 * time.Second

	client, err := server.Connect()
	if err != nil {
		log.Fatal("[package main]:[func SendMessage] - cannot create client")
	}

	// Creating an empty email message, and giving an info to it
	email := mail.NewMSG()
	email.SetFrom(mailData.From).AddTo(mailData.To).SetSubject(mailData.Subject)
	email.SetBody(mail.TextHTML, mailData.Content)

	err = email.Send(client)
	if err != nil {
		log.Fatal("[package main]:[func SendMessage] - cannot sent email")
	}
}
