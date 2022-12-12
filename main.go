package main

import (
	"fmt"
	"Golang_Struct/contas"
	//"Golang_Struct/clientes"

)

func PagarBoleto(conta verificarConta, valorDoBoleto float64){
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaAnderson := contas.ContaPoupanca{}
	contaAnderson.Depositar(300)

	fmt.Println(contaAnderson)

	contaAnderson.Sacar(100.)

	fmt.Println(contaAnderson)
	
	PagarBoleto(&contaAnderson, 120)

	fmt.Println(contaAnderson.ObterSaldo())

	contaAna := contas.ContaCorrente{}
	contaAna.Depositar(400)

	fmt.Println(contaAna)

	contaAna.Sacar(100)

	fmt.Println(contaAna)

	PagarBoleto(&contaAna, 120)

	fmt.Println(contaAna.ObterSaldo())
	}