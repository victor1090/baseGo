#Slice

##Pecorrendo slice
slice := []int{1, 2, 3, 4}
for indice, valor := range slice {
	fmt.Printf("O indice %v tem o valor %v. \n", indice, valor)
}

##Fatiando slice
fatia:= slice[0:2]
fatia = {1,2}
###Indo ate o final
fatia:= slice[1:]

##Deletando valor 3 do slice
slice = append(slice[:2], slice[3:]...)


##Para performance de Slice é possível alocar memoria antecipadamente usando make
slice := make([]int, 5 ,10)  onde 5 é o len do slice e 10 é a capacity do array.
funções para ver: len(slice), cap(slice)


##Slice Multidimensional
ss:= [][]int{ []int{1,2,3}, []int{4,5,6}, []int{7,8,9},}


#Maps

##Definição
amigos := map[string]int{
	"victor": 99391381,
	"joana": 99138166,
}

##Pra adicionar basta
amigos["lucas"] = 4444444

##Pra saber se o valor 0 é o valor dentro do map ou null basta fazer
valor,ok := amigos["naoExiste"] resultando em 0 false

##Pecorrendo Maps( range não garante a ordem)
for key, value :=range amigos{
	fmt.Println(key, value)
}

##Deletando do Maps
delete(amigos,key)
