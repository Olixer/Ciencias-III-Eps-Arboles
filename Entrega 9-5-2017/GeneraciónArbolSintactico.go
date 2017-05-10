package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//Estructura del arbol y sus nodos

type Arbol struct {
	Izquierda *Arbol
	Valor     string
	Derecha   *Arbol
}

func NewArbol() *Arbol {
	return &Arbol{}
}

type Stack struct {
	arbol []*Arbol
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(n *Arbol) {
	s.arbol = append(s.arbol[:s.count], n)
	s.count++
}

func (s *Stack) Pop() *Arbol {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.arbol[s.count]
}

type Variables struct {
	Ecuacion         string
	Valor            string
	Variable         string
	EcuacionOriginal string
}

type StackVariables struct {
	stackV []*Variables
	count  int
}

func NewStackVariables() *StackVariables {
	return &StackVariables{}
}

func (s *StackVariables) PushV(n *Variables) {
	s.stackV = append(s.stackV[:s.count], n)
	s.count++
}

func (s *StackVariables) PopV() *Variables {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.stackV[s.count]
}

//Funciones que afectan el ingreso de ecuaciones

//Funcion que hace uso de la ecu ingresada para hacer uso de los valores numericos
func encontrarVariable(cadenaCompleta string) ([]string, []string, string) {
	s1 := cadenaCompleta
	if last := len(s1) - 1; last >= 0 && s1[last] == '=' {
		s1 = s1[:last]
	}
	if last := len(s1) - 1; last >= 0 && s1[last] == ' ' {
		s1 = s1[:last]
	}
	if last := len(s1) - 1; last >= 0 && s1[last] == ':' {
		s1 = s1[:last]
	}
	if last := len(s1) - 1; last >= 0 && s1[last] == ' ' {
		s1 = s1[:last]
	}
	arr := strings.Split(s1, " ")
	varbls := []string{}
	cadenaTotal := []string{}
	variableNueva := arr[len(arr)-1]
	arr = arr[:len(arr)-1]
	for i := 0; i < len(arr); i++ {
		if arr[i] != "+" && arr[i] != "-" && arr[i] != "*" && arr[i] != "/" {
			if _, err := strconv.Atoi(arr[i]); err == nil {
			} else {
				varbls = append(varbls, arr[i])
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		cadenaTotal = append(cadenaTotal, arr[i])
	}
	return cadenaTotal, varbls, variableNueva
}

//Ingresar el valor de las varbls segun su ecuación
func (s *StackVariables) ValorVar(names []string) []string {
	num := []string{}
	for i := 0; i < len(names); i++ {
		for j := 0; j < s.count; j++ {
			if s.stackV[j].Variable == names[i] {
				num = append(num, s.stackV[j].Valor)
			}
		}
	}
	return num
}

func ordenarEcu(ecu []string, valores []string, varbls []string) string {
	finalEcu := ""
	for i := 0; i < len(ecu); i++ {
		finalEcu += ecu[i]
		finalEcu += " "
	}
	if final := len(finalEcu) - 1; final >= 0 && finalEcu[final] == ' ' {
		finalEcu = finalEcu[:final]
	}
	result := finalEcu
	for i := 0; i < len(valores); i++ {
		result = strings.Replace(result, varbls[i], valores[i], -1)
	}
	return result
}

func insertPila(a string) *Arbol {
	stack := NewStack()
	val := strings.Split(a, " ")
	for i := 0; i < len(val); i++ {
		arbol := &Arbol{nil, val[i], nil}
		if arbol.Valor != "+" && arbol.Valor != "-" && arbol.Valor != "*" && arbol.Valor != "/" {
			stack.Push(arbol)
		} else {
			arbol.Derecha = stack.Pop()
			arbol.Izquierda = stack.Pop()
			stack.Push(arbol)
		}
	}
	return stack.Pop()
}

func Operacion(a *Arbol) int {
	if a == nil {
		return 0
	} else if a.Valor == "+" {
		return Operacion(a.Izquierda) + Operacion(a.Derecha)
	} else if a.Valor == "-" {
		return Operacion(a.Izquierda) - Operacion(a.Derecha)
	} else if a.Valor == "*" {
		return Operacion(a.Izquierda) * Operacion(a.Derecha)
	} else if a.Valor == "/" {
		if Operacion(a.Derecha) == 0 {
			fmt.Println("\nGO Error 4 Division por cero")
		} else {
			return Operacion(a.Izquierda) / Operacion(a.Derecha)
		}
		return 0
	} else {
		conv, _ := strconv.Atoi(a.Valor)
		return conv
	}
}

func quitaVariable(v string) string {
	s1 := v
	if last := len(s1) - 1; last >= 0 && s1[last] == '=' {
		s1 = s1[:last]
	}
	if last := len(s1) - 1; last >= 0 && s1[last] == ' ' {
		s1 = s1[:last]
	}
	if last := len(s1) - 1; last >= 0 && s1[last] == ':' {
		s1 = s1[:last]
	}
	if last := len(s1) - 1; last >= 0 && s1[last] == ' ' {
		s1 = s1[:last]
	}
	arr := strings.Split(s1, " ")
	arr = arr[:len(arr)-1]

	cadena := strings.Join(arr, " ")
	return cadena
}

func (s *StackVariables) imprimirVariable() string {
	return s.stackV[s.count-1].Variable
}

//Verificar si la variable cumple con los criterios
func (s *StackVariables) varValida() bool {
	v := s.stackV[s.count-1].Variable

	var validar = regexp.MustCompile(`^[A-Za-z]+$`)
	verificar := validar.MatchString(v)

	if verificar == true {
		return true
	}
	return false
}

//Verificar si la variable cumple con los criterios
func validarVariable(v string) bool {

	var validar = regexp.MustCompile(`^[A-Za-z]+$`)
	verificar := validar.MatchString(v)

	if verificar == true {
		return true
	}
	return false
}

//Funciones que afectan la clasificacion de tipos en las ecuaciones

func (s *StackVariables) tiposVariable() {
	for i := 0; i < s.count; i++ {
		ecu := s.stackV[i].Ecuacion
		arr := strings.Split(ecu, " ")
		for j := 0; j < len(arr); j++ {
			var valorValido = regexp.MustCompile(`^[0-9]+$`)
			if valorValido.MatchString(arr[j]) == true {
				fmt.Println("Valor: ", arr[j])
			}

			if arr[j] == "+" || arr[j] == "-" || arr[j] == "*" || arr[j] == "/" {
				fmt.Println("Operador: ", arr[j])
			}
		}

		fmt.Println(fmt.Sprint("Identificador: ", s.stackV[i].Variable))
		fmt.Println("Operador: :=")
		fmt.Println("\n")
	}
}

//Funciones que afectan las ecuaciones ingresadas por consola

func (s *StackVariables) formIngresada() {
	for i := 0; i < s.count; i++ {

		ecuacionOriginal := s.stackV[i].EcuacionOriginal
		arrOriginal := strings.Split(ecuacionOriginal, " ")

		fmt.Println("Ecuacion: ", ecuacionOriginal)

		for m := 0; m < len(arrOriginal); m++ {
			if validarVariable(arrOriginal[m]) == true {
				for t := 0; t < s.count; t++ {
					if arrOriginal[m] == s.stackV[t].Variable {
						arrOriginal[m] = s.stackV[t].EcuacionOriginal
					}
				}
			}
		}
		cadenaOriginal := strings.Join(arrOriginal, " ")
		fmt.Println("Ecuacion sin identificadores: ", cadenaOriginal)
		result := insertPila(cadenaOriginal)
		fmt.Println("Ecuacion en infijo: ")
		inorden(result)
		fmt.Println("\nResultado = ", Operacion(result), "\n")
	}
}

func inorden(a *Arbol) {
	if a == nil {
		return
	}
	inorden(a.Izquierda)
	fmt.Printf(a.Valor)
	inorden(a.Derecha)
}

func menu(s *StackVariables) {
	var opcion int
	leer := bufio.NewScanner(os.Stdin)
	leer.Scan()
	fmt.Println("\n")
	opcionIngresada := leer.Text()
	opcion, fallo := strconv.Atoi(opcionIngresada)
	if fallo != nil {
		fmt.Println(fmt.Sprint("GO Error 1 Tipo de dato"))
	}
	switch opcion {
	case 1:
		fmt.Println("EJ: 7 2 + 8 - X :=")
		fmt.Println("\nIngresar ecuación en postfijo\n")
		leer := bufio.NewScanner(os.Stdin)

		for leer.Scan() {

			opcionIngresada := leer.Text()
			ecuacionAux, varbls, variableAux := encontrarVariable(opcionIngresada)
			ecu := s.ValorVar(varbls)
			finalEcu := ordenarEcu(ecuacionAux, ecu, varbls)
			resultado := insertPila(finalEcu)
			FinalV := strconv.Itoa(Operacion(resultado))
			ecuNoVarbl := quitaVariable(opcionIngresada)
			s.PushV(&Variables{finalEcu, FinalV, variableAux, ecuNoVarbl})

			if s.varValida() == false {
				fmt.Println("\nGO Error 2 Tipo variable invalido")
				os.Exit(3)
			}
			break
		}
	case 2:
		s.tiposVariable()
		fmt.Println("\n")
	case 3:
		s.formIngresada()
		fmt.Println("\n")
	case 4:
		os.Exit(3)
		fmt.Println("\n")
	default:
		fmt.Println("GO Error 3 Opcion invalida")
	}
}

func main() {
	stack := NewStackVariables()

	for {
		fmt.Println("1. Ingresar ecuaciones ")
		fmt.Println("2. Tipos de varbls")
		fmt.Println("3. Formulas ingresadas ")
		fmt.Println("4. Exit\n")
		fmt.Print("Opcion: ")
		menu(stack)

	}

}
