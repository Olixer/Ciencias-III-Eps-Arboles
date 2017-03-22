package main

import (
	bu "bufio"
	io "fmt"
	os "os"
	conv "strconv"
	str "strings"
	regexp "regexp"
)


type Nodo struct {
	Valor  int
	Nombre string
}

type Stack struct {
	nodos    []*Nodo
	contador int
}

func RecibirDatos() (string, error) {
	leer := bu.NewReader(os.Stdin)
	s, err := leer.ReadString('\n')

	return str.TrimSpace(s), err
}

func (nodo *Nodo) String() string {
	return nodo.Nombre
}

func NuevoStack() *Stack {
	return &Stack{}
}

func (pila *Stack) Push(nodo *Nodo) {
	pila.nodos = append(pila.nodos[:pila.contador], nodo)
	pila.contador++
}

func (pila *Stack) Pop() *Nodo {
	if pila.contador == 0 {
		return nil
	}
	pila.contador--
	return pila.nodos[pila.contador]
}

func ResolverPila(pila *Stack) int {
	pilaAux := NuevoStack()
	rsta := 0
	for i := 0; i < len(pila.nodos); i++ {
		termino := pila.Pop()
		aux, err := conv.Atoi(termino.Nombre)
		if err != nil {
			switch termino.Nombre {
			case "+":
				rsta = pilaAux.Pop().Valor + pilaAux.Pop().Valor
				pilaAux.Push(&Nodo{rsta, ""})
			case "-":
				rsta = pilaAux.Pop().Valor - pilaAux.Pop().Valor
				pilaAux.Push(&Nodo{rsta, ""})
			case "/":
				denominador := pilaAux.Pop().Valor
				if denominador != 0 {
					rsta = pilaAux.Pop().Valor / denominador
				} else {
					rsta = pilaAux.Pop().Valor / 1
				}
				pilaAux.Push(&Nodo{rsta, ""})
			case "*":
				rsta = pilaAux.Pop().Valor * pilaAux.Pop().Valor
				pilaAux.Push(&Nodo{rsta, ""})
			}
		} else {
			pilaAux.Push(&Nodo{aux, ""})
		}
	}
	return rsta
}

func main() {
	errorExpresion := 0
	io.Println("Ejemplo: Palabra_00 := + 5 3")
	io.Print("Ingrese la expresion en postfijo: ")
	arbol1String, err := RecibirDatos()
	pila1 := NuevoStack()
	array1 := str.Split(arbol1String, " ")
	variable, _ := regexp.Compile("^[A-Z]([a-zA-Z0-9]+)_*([0-9])")
    if(variable.MatchString(array1[0])==false){
		errorExpresion=errorExpresion+1
	}
	resultado, _ := regexp.Compile(":=")
    if(resultado.MatchString(array1[1])==false){
		errorExpresion=errorExpresion+1
	}
	if(errorExpresion==0){
		for i := 0; i < len(array1); i++ {
			pila1.Push(&Nodo{i, array1[i]})
		}

		x := ResolverPila(pila1)
		arr := [2]int{x}

		io.Println(array1[0],"=", arr[0])

		if err != nil {
			io.Println("Error cuando se escaneo la entrada", err.Error())
			return
		}
	}else{
		io.Print("Expresion erronea")
	}

}
