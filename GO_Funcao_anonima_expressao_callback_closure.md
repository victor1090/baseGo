#Função Anonima
##Função executada uma unica vez(descartavel) não precisa de nome e o argumento é passado no final
func main(){
  func(x int) {
      fmt.Println(x, "vezes 2 é:")
      fmt.Println(x * 2)
  }(5)
}

#Função como Expressão
##Atribui a função a uma variavel

func main(){
  x:= 5
  y:= func(x int) {
        return x * 2
      }
  fmt.Println(x, "vezes 2 é:", y(x))
}

#Retornando uma Função

func retornaf(x int) func(int) int {
  return func(i int ) int {
      return i * 10
  }
}

##usando
func main(){
  x:= retornaf()
  y:= x(3)
  fmt.Println(y)

  // Pode ser chamada dessa forma sem o x como intermediario
  fmt.Println(retornaf()(3))
}

#Callback
## Um callback em go é feito passando uma função como argumento de outra função para ser executada dentro por exemplo:

func soma(x ...int) int {
    sum:=0
    for _, v := range x{
        sum += v
    }
    return sum
}

##O callback vem agora
func somaPares(nomefunc func(x ...int) int, y ...int) {
    var slice []int
    for _, v := range y{
        if v % 2 == 0 {
            slice = append(slice, v)
        }
    }
    //Agora vamos executar a chamada de callback. chamando a função recebida
    // no argumento apenas com o slice que contem os números pares e entao ele vai somar
    total := nomefunc(slice...)
    return total
}

func main(){
  t:= somaPares(soma,[]int{1,2,3,4,5,6,7,8,9,10}...)
  fmtPrintln(t)
}

#Closure
##Closure significa capturar o contexto, por exemplo uma função tem um contexto, um package tem um contexto, indentifica
## No Closure capturamos um contexto e criamos uma copia dele para usar em outra lugar veja o exemplo:

func i() func() int {
    x:= 0
    return func() int{
        x++
        return x
    }
}

##Perceba que a variavel x não faz parte do escopo da função interna que vai ser retornada, mas capturamos o valor de x e guardamos para utilizar dentro da função retornada
##Caso i() seja chamada duas vezes para variaveis diferentes o valor de x contido no escopo interno será diferente pois o closure cria uma copia do valor para usar dentro
#Lembrando que GO é pass by value e não pass by reference
