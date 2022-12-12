package main

import (
	"fmt"
	"Golang_Struct/contas"
	//"Golang_Struct/clientes"

)

func main() {
	contaAnderson := contas.ContaPoupanca{}
	contaAnderson.Depositar(100)

	fmt.Println(contaAnderson)

	contaAnderson.Depositar(200)

	fmt.Println(contaAnderson)

	contaAnderson.Sacar(50.)

	fmt.Println(contaAnderson.ObterSaldo())

} 