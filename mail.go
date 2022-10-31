package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

/*
	func SendMail(cfg Config) {
		var tpl bytes.Buffer
		t, err := template.ParseFiles("data/tmplExpiry.html")
		if err != nil {
			log.Println(err)
		}

		if err := t.Execute(&tpl, license); err != nil {
			log.Println(err)
		}

		manager := []string{license.Owner}
		admins := cfg.Dictionaries.AdminsEmails
		sendMail(cfg.SMTP.Host, cfg.SMTP.Port, cfg.SMTP.Username, cfg.SMTP.Password, cfg.SMTP.From, manager, admins, "BM license expires", tpl.String(), "")
	}
*/
func SendMail(host string, port int, user string, password string, from string, to string, cc string, subject string, body string, attach string) {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Cc", cc)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if attach != "" {
		m.Attach(attach)
	}

	d := gomail.NewDialer(host, port, user, password)

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
	}
}
