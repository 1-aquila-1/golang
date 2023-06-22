package gmail_test

import (
	"testing"

	"github.com/1-aquila-1/golang/gmail"
)

func TestEnviarEmailSimple(t *testing.T) {
	if erro := gmail.EmailSimple(); erro != nil {
		t.Fatal(erro)
	}
}

func TestEnviarEmailComTemplate(t *testing.T) {
	if erro := gmail.EmailComTemplate("./template.html"); erro != nil {
		t.Fatal(erro)
	}
}
