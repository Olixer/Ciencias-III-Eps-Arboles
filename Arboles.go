package main
import (
  "fmt"
  "strconv"
)
type ArbolExpresion struct{
  Izquierda *ArbolExpresion
  Valor string
  Derecha *ArbolExpresion
}
func Calcular(t *ArbolExpresion) int{
  if t.Valor=="+"{
    return  Calcular(t.Izquierda) + Calcular(t.Derecha)
  }
  if t.Valor=="-"{
    return Calcular(t.Izquierda) - Calcular(t.Derecha)
  }
  if t.Valor=="*"{
    return Calcular(t.Izquierda) * Calcular(t.Derecha)
  }
  if t.Valor=="/"{
    return Calcular(t.Izquierda) / Calcular(t.Derecha)
  }
    i:=strconv.Atoi(t.Valor)
    return i

}
func Expresar(t *ArbolExpresion) string{
  if t.Valor=="+"{
    return  Expresar(t.Izquierda) +" + "+ Expresar(t.Derecha)
  }
  if t.Valor=="-"{
    return Expresar(t.Izquierda) +" - "+ Expresar(t.Derecha)
  }
  if t.Valor=="*"{
    return Expresar(t.Izquierda)+" * "+Expresar(t.Derecha)
  }
  if t.Valor=="/"{
    return Expresar(t.Izquierda)+" / "+ Expresar(t.Derecha)
  }
    i:=(t.Valor)
    return i
}
func main(){
  t1:= &ArbolExpresion{&ArbolExpresion{&ArbolExpresion{nil,"9", nil},"*",&ArbolExpresion{nil,"3",nil}},"+",&ArbolExpresion{&ArbolExpresion{nil,"7",nil},"-",&ArbolExpresion{nil,"4",nil}}}
  fmt.Println(Calcular(t1))
  fmt.Println(Expresar(t1))
}

