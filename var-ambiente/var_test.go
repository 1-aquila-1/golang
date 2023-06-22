package varambiente_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestBuscandoValorDaVariavel(t *testing.T) {
	if erro := godotenv.Load(".env"); erro != nil {
		t.Fatal(erro)
	}
	fmt.Println(os.Getenv("EMAIL_SENHA"))
}

func TestRetornandoVariaveisArquivo(t *testing.T) {
	data, erro := godotenv.Read(".env")
	if erro != nil {
		t.Fatal(erro)
	}
	fmt.Println(data["EMAIL_SENHA"])
}
