package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"text/template"
)

type EmailMensagem struct {
	Mensagem string
}

func emailSimple(auth smtp.Auth) {
	msg := "Subject: Primeiro e-mail\nUm e-mail de teste 23"
	erro := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"teste.teste4@gmail.com",
		[]string{"teste.teste4@gmail.com"},
		[]byte(msg),
	)
	if erro != nil {
		fmt.Println(erro)
	}
}

func emailHtml(auth smtp.Auth, templatePath string) {

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)

	if err != nil {
		log.Fatalln(err)
	}

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	t.Execute(&body, EmailMensagem{Mensagem: "Aparecem flores na terra, e chegou o tempo de cantar; já se ouve em nossa terra o arrulhar dos pombos. Cânticos 2:12"})

	msg := "Subject: Versículo de Hoje\n" + headers + "\n\n" + body.String()
	erro := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"teste.teste4@gmail.com",
		[]string{"teste3.desenv@gmail.com"},
		[]byte(msg),
	)
	if erro != nil {
		fmt.Println(erro)
	}
}

func main() {
	auth := smtp.PlainAuth(
		"",
		"teste.teste@gmail.com",
		"senha_email",
		"smtp.gmail.com",
	)

	emailHtml(auth, "./template.html")

}
