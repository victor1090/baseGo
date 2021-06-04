package main

import (
	"fmt"
)

//Com slice
func menorMaior(x []int) (int, int) {
	menor:= x[0]
	maior:= x[0]
	for _, v := range x{
		if(menor > v){menor = v}
		if(maior < v){maior = v}
	}
	return menor, maior
}
//Com Parametro Variatico
func menorMaior2(x ...int) (int, int) {
	menor:= x[0]
	maior:= x[0]
	for _, v := range x{
		if(menor > v){menor = v}
		if(maior < v){maior = v}
	}
	return menor, maior
}

func main() {
	x := []int{2,3,4,5,6,7}
	menor, maior:=menorMaior(x)
	menor2,maior2:=menorMaior2(2,3,4,5,6,7)
	//Também é possivel desenrolar a slice para enviar item por item para a funão 2 por exemplo
	menor3,maior3:=menorMaior2(x...)
	fmt.Println(menor,maior)
	fmt.Println(menor2,maior2)
	fmt.Println(menor3,maior3)
}
