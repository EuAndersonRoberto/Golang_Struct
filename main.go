package main

import (
	"fmt"
	"Golang_Struct/contas"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contadoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}
	
	status := contaDaSilvia.Transferir(200, &contadoGustavo)
	
	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contadoGustavo)
}