package main

import (
	"encoding/json"
	"gopkg.in/gomail.v2"
	"log/slog"
)

var emails = []string{"sads1100@gmail.com"}

func SendMessage(post Post) {
	dialer := gomail.NewDialer("smtp.mail.ru", 465, "aalex110@mail.ru", "rRUKsTdspd5njv3AK1fJ")

	body, _ := json.Marshal(post)

	for _, email := range emails {
		message := gomail.NewMessage()
		message.SetHeader("From", "aalex110@mail.ru")
		message.SetHeader("To", email)
		message.SetHeader("Subject", "НОВЫЙ ПОСТ У ТВОЕГО ХЗ КОГО")
		message.SetBody("text/html", string(body))
		if err := dialer.DialAndSend(); err != nil {
			slog.Error(err.Error())
		}
	}

}
