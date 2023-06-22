package gmail

import (
	"bytes"
	"html/template"
	"net/smtp"
)

var auth = smtp.PlainAuth(
	"",
	"aquilla.tavares7@gmail.com",
	"wdbdwsfqxggwwyyu",
	"smtp.gmail.com",
)

var objMsg = EmailMensagem{
	Mensagem: "Aparecem flores na terra, e chegou o tempo de cantar; já se ouve em nossa terra o arrulhar dos pombos. Cânticos 2:12",
}

func EmailSimple() error {
	msg := "Subject: Primeiro e-mail\nUm e-mail de teste 1"
	erro := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"aquilla.tavares7@gmail.com",
		[]string{"aquila.desenv@gmail.com"},
		[]byte(msg),
	)
	return erro
}

func EmailComTemplate(templatePath string) error {
	var body bytes.Buffer
	t, erro := template.ParseFiles(templatePath)
	if erro != nil {
		return erro
	}
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	erro = t.Execute(&body, objMsg)
	if erro != nil {
		return erro
	}

	msg := "From:Rede Point das Baterias<aquilla.tavares7@gmail.com>\n"
	msg += "To: aquila.desenv@gmail.com\n"
	msg += "Subject: Versículo de Hoje\n" + headers + "\n" + body.String()
	to := []string{"aquila.desenv@gmail.com"}
	erro = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"aquilla.tavares7@gmail.com",
		to,
		[]byte(msg),
	)
	return erro
}
