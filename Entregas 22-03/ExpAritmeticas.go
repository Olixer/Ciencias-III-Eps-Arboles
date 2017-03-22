package main

import(
	io "fmt"
	str "strings"
	conv "strconv"
)



type Nodo struct{
	Valor int
	Nombre string
}

type Stack struct{
	nodos []*Nodo
	contador int
}



func (nodo *Nodo) String() string{
	return nodo.Nombre
}


func NuevoStack() *Stack {
	return &Stack{}
}


func (pila *Stack) Pop() *Nodo{
	if pila.contador == 0{
		return nil
	}
	pila.contador--
	return pila.nodos[pila.contador]
}

func (pila *Stack) Push(nodo *Nodo){
	pila.nodos = append(pila.nodos[:pila.contador], nodo)
	pila.contador++
}



func ResolverPilas(pila *Stack) int{
	pilaAux := NuevoStack()
	rsta := 0
	for i:=0; i<len(pila.nodos); i++{		
		termino := pila.Pop()
		aux, err := conv.Atoi(termino.Nombre)
		if err != nil{
			switch termino.Nombre{
				case "+":					
					rsta = pilaAux.Pop().Valor + pilaAux.Pop().Valor
					pilaAux.Push(&Nodo{rsta,""})
				case "-":
					rsta = pilaAux.Pop().Valor - pilaAux.Pop().Valor
					pilaAux.Push(&Nodo{rsta,""})
				case "/":
					denominador := pilaAux.Pop().Valor
					if denominador != 0 {
						rsta = pilaAux.Pop().Valor / denominador						
					}else{
						rsta = pilaAux.Pop().Valor / 1
					}
					pilaAux.Push(&Nodo{rsta,""})
				case "*":
					rsta = pilaAux.Pop().Valor * pilaAux.Pop().Valor
					pilaAux.Push(&Nodo{rsta,""})
			}
		}else{
			pilaAux.Push(&Nodo{aux,""})
		}
	}
	return rsta
}

func main(){
	
	pila1 := NuevoStack()
	pila2 := NuevoStack()
	pila3 := NuevoStack()
    arboluno := ":= + 5 3 x"
	arboldos := ":= + - 5 2 4 y"
	arboltres := ":= / x y z"
	array1 := str.Split(arboluno, " ")
	array2 := str.Split(arboldos, " ")
	array3 := str.Split(arboltres, " ")
	for i:=0; i<len(array1); i++{
		pila1.Push(&Nodo{i,array1[i]})
	}
	for i:=0; i<len(array2); i++{
		pila2.Push(&Nodo{i,array2[i]})
	}
	for i:=0; i<len(array3); i++{
		pila3.Push(&Nodo{i,array3[i]})
	}
	x := ResolverPilas(pila1)
	y := ResolverPilas(pila2)
	arr := [2] int {x,y}
	for i:=0; i<=len(arr); i++{
		pila3.Pop()
	}
	for i:=0; i<len(arr); i++{
		pila3.Push(&Nodo{i,conv.Itoa(arr[i])})
	}
	z  := ResolverPilas(pila3)
	io.Println("X=",arr[0]," Y=",arr[1]," Z=",z)

}
