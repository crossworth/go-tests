package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// Redeclarar variável
	// somente é possível caso uma das variáveis seja nova
	// ou caso o escopo seja diferente
	i, b := 10, 20
	i, c := 200, 20

	_ = b
	_ = c

	fmt.Println(i)

	var str = "teste de stringç"

	for i, c := range str {
		fmt.Printf("I: %v %c %p \n", i, c, &c)
	}

	var t interface{} = "teste"

	fmt.Println(t.(string)) // assert que é uma string
	asInt, ok := t.(int)    // assert que é um int
	fmt.Println(asInt, ok)

	var i2 interface{} = 10

	if _, isInt := i2.(int); isInt {
		fmt.Println("Is int")
	} else {
		fmt.Println("Is not int")
	}

	// new aloca a memoria zerada
	pointer := new([1024 * 1024 * 1024]int)
	_ = pointer

	// time.Sleep(20 * time.Second)

	// no caso de structs
	type mys struct {
		n int
		s uint64
	}

	pm := new(mys)
	_ = pm
	// é a mesma coisa que fazer isso
	pm = &mys{}

	// make é utilizado para inicializar slices, maps e channels
	// make não retorna pointer igual new

	m := make(map[string]bool, 0) // type, size, [capacidade]
	m = map[string]bool{}         // igual a linha acima
	m = map[string]bool{          // inicialização com valores
		"test":    true,
		"working": false,
	}
	_ = m

	// o make no caso de um slice recebe type, length, capacity
	// um slice possui na verdade é composto por um array
	// quando passamos um capacity, estamos definindo o valor que
	// esperamos guardar dentro desse array (o tamanho do array interno)
	// porém caso utilizemos append, mesmo estourando a capacidade,
	// internamente o append realoca um novo array
	s := make([]int, 10, 11) // retorna um array com 10 elementos zerados
	fmt.Printf("Len %v Cap %v\n", len(s), cap(s))
	s = append(s, 1) // adiciona 1 elemento
	s = append(s, 1) // adiciona 1 elemento e dobra a capacidade do array
	fmt.Println(s)
	fmt.Printf("Len %v Cap %v\n", len(s), cap(s))

	// Slice of slice
	mm := make([][]string, 0)
	fmt.Println(mm)
	outer := make([]string, 0)
	outer = append(outer, "teste1")
	outer = append(outer, "teste2")
	mm = append(mm, outer)
	fmt.Println(mm)

	// maps
	var nomesIdade = map[string]int{
		"Pedro": 25,
		"Teste": 10,
	}

	fmt.Println(nomesIdade)
	fmt.Println(nomesIdade["Pedro"])

	// check if a value exists, by default map return empty when not
	// present

	if _, ok := nomesIdade["pedroh"]; !ok {
		fmt.Println("Não encontrado")
	}

	// delete item
	delete(nomesIdade, "Pedro")

	re := refc{0}
	re.MethodOne()

	// zero, temos metodos no pointer
	fmt.Println("Num of methods", reflect.TypeOf(re).NumMethod())
	fmt.Println("Num of methods pointer", reflect.TypeOf(&re).NumMethod())

	ma := map[string]string{
		"test": "123",
	}

	// podemos utilizar o map acessando diretamente o chave (ele irá retornar um valor vazio)
	// se a mesma não existir, ou podemos usar uma segunda variavel para verificar se a chave existe
	rMap, ok := ma["test"]
	fmt.Println(rMap)

	// converter int to string
	fmt.Println(strconv.Itoa(10))

}

type refc struct {
	tng int
}

func (t *refc) MethodOne() {
	t.tng = 10
}

func (t *refc) MethodTwo() {
	t.tng = 20
}

func (t *refc) MethodThree() {
	t.tng = 30
}
