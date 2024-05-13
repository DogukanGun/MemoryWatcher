package utils

import (
	"log"
	"net/smtp"
	"os"
)

func Send(body string) {
	from := "freelance.dogukang@gmail.com"
	pass := os.Getenv("EMAIL_PASSWORD")
	to := "dogukangundogan5@gmail.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Println("Successfully sended to " + to)
}
