package main

import (
	"fmt"

	"github.com/1-aquila-1/golang/poo/conta"
)

func main() {
	contaDaSilvia := conta.ContaCorrente{Titular: "Silvia", Saldo: 500}
	fmt.Println(contaDaSilvia.Saldo)
	fmt.Println(contaDaSilvia.Sacar(300))
	fmt.Println(contaDaSilvia.Saldo)
}
