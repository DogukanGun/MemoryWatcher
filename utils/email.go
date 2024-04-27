package utils

import "net/smtp"

// Function to send email notification
func sendEmail(subject, body string) error {
	from := "your-email@example.com"
	password := "your-email-password"
	to := "your-email@example.com"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail("smtp.example.com:587",
		smtp.PlainAuth("", from, password, "smtp.example.com"),
		from, []string{to}, []byte(msg))

	return err
}
