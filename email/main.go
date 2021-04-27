package main

import "github.com/go-gomail/gomail"

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "<TODO>@gmail.com")
	// m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Omar</b>!")
	// m.Attach("file.png")

	// TODO: find smtp server that works
	d := gomail.NewDialer("smtp.gmail.com", 587, "username", "password")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

// http://stackoverflow.com/questions/34790771/how-do-i-insert-an-image-into-email-body
