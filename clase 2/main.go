package main

import (
	"fmt"
)
	
//defer y panic
func main() {
	miFuncion()

}

func miFuncion() {
	defer fmt.Println("mensaje final")
	fmt.Println("inicio de la funcion")
	a := 10
	if a==10 {
		panic("fallo")
	}
	fmt.Println("fin de la funcion")
}
/*

//recursividad
func main() {
	miFuncion(1)

}

func miFuncion(valor int) {
	dato:= valor +1
	fmt.Println("el valor es:", dato)
	if dato<10 {
		miFuncion(dato)
	}
}

//gorutinas
func main() {
	fmt.Println("inicio del programa")
	time.Sleep(5 * time.Second)
	fmt.Println( miFuncion("Gorutinas") )
	//ejemplo
	miCanal := make (chan string)
	go func() {
		time.Sleep(3 * time.Second)
		miCanal <- "mensaje desde la gorutina"
	}()
	fmt.Println(<- miCanal)
	fmt.Println("fin del programa")
}

func miFuncion(parametro string) string{
	return "hola mundo " + parametro	
}

//Ejercicio: Funciones con y sin parámetros
func main() {
miFuncion()
miFuncionConParametros(10, 20)
fmt.Println( miFuncionConRetorno("Juan"))
fmt.Println( miFuncionretornoMultiple(15, 5) )
fmt.Println("La suma es:", suma(5, 7) )
//closure
tabla := tabla(2)
for i:= 1; i <= 10; i++ {
	fmt.Println(tabla())
}
}

func miFuncion() {
	fmt.Println("inicio de la funcion")
}

func miFuncionConParametros(numero1 int, numero2 int) {
	resultado := numero1 + numero2
	fmt.Println("La suma es :", resultado)
}

func miFuncionConRetorno(nombre string) string {
	return "Hola " + nombre
}

func miFuncionretornoMultiple(numero1 int, numero2 int) (int, int) {
	suma := numero1 + numero2
	resta := numero1 - numero2
	return suma, resta
}

var suma = func (numero1 int, numero2 int) int  {
	return numero1 + numero2	
}

//closure
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}




*/