package utils

import (
	"fmt"
	"net/smtp"
)

func SendUpdatePassword(url *string) {

	from := "wijayakusumasandi@gmail.com"
	password := "cmtBOUD78JAF6NYb"

	// Receiver email address.
	to := []string{
		"kristono.bricks@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp-relay.brevo.com"
	smtpPort := "587"

	// Message.
	message := []byte("URLPARSE")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("SEND_EMAIL_ERROR", err)
		return
	}
	fmt.Println("Email Sent Successfully!")

}
