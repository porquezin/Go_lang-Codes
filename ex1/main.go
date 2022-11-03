package main

import (
	"ex1/contas"
	"fmt"
)

type verificarConta interface {
	Sacar(saque float64) (string, float64)
}

func PagarBoleto(conta verificarConta, boleto float64) {
	conta.Sacar(boleto)
}

func main() {
	conta1 := contas.ContaPoupanca{}
	conta1.Depositar(100)
	PagarBoleto(&conta1, 60)
	fmt.Println(conta1.GetSaldo())
}
