#Struct
##Definição
type cliente struct {
    nome        string
    sobrenome   string
    fumante     bool
}


cliente1 := cliente{
              nome:       "João",
              sobrenome:  "da Silva",
              fumante:    false,
}

cliente2:= cliente{"Joana", "Pereira", true}

##Structs Composto
Se atentar para o tipo pessoa dentro de profissional na declaracao fica pessoa: pessoa
pessoa1:= profissional{
    pessoa: pessoa{
      nome:   "Maria",
      idade:  31,
    },
    título:   "Engenheira",
    salario:  5000,
}

#Acessando os campos
pessoa1.titulo
pessoa1.pessoa.nome ou pesso1.nome caso não tenha conflito

#Struct Anonimo
Declarando um struct no main e usando
x:= struct{
  nome string
  idade int
}{
  nome: "Mario",
  idade: 50
}
