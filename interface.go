package main

import (
	"fmt"
)

type pessoa struct {
	nome 	  string
	sobrenome string
	idade 	  int
}

type dentista struct {
	pessoa
	especializacao 	string
	salario 	int
}

type engenheiro struct {
	pessoa
	salario	int
	empresa	string
}

func (p dentista) bemVindo(){
	fmt.Println(p.nome, "diz bem vindo ao meu consultorio!")
}

func (p engenheiro) bemVindo(){
	fmt.Println(p.nome, "diz bem vindo a obra escravo!")
}

// Essa interface possue o metodo bemVindo e como engenheiro e dentista implementam
// esse metodo eles fazem parte dessa interface. Na hora da chamada a interface
// indentifica quem ta chamando e ou chama as funções de cada uma dela ou chama metodos especificos.
// Você pode ter varios objetos diferentes mas no código irá executá-los da mesma forma.
type humano interface {
	bemVindo()
}

func apresentacao(h humano) {
	h.bemVindo()
	switch h.(type) {
		case dentista:
			fmt.Println("Eu sou especialista em", h.(dentista).especializacao)
		case engenheiro:
			fmt.Println("Eu trabalho na empresa", h.(engenheiro).empresa)
	}
}

func main() {
	pedro := dentista{pessoa{"Pedro","Silva",30},"Cirurgiao",5000}
	victor := engenheiro{pessoa{"Victor", "Santos",30},5000,"Labsit"}
	//victor.bemVindo()
	//pedro.bemVindo()
	apresentacao(pedro)
	apresentacao(victor)
}
