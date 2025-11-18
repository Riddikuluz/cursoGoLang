package main

import (
	"fmt"
)

//interfaces
func main() {	
	e := Estructuras{}
	fmt.Println(e.Mifuncion())
	fmt.Println(e.calculo(5, 10))
}

type Ejemplo interface {
	Mifuncion() string
	calculo(n1 int, n2 int) int
}

type Estructuras struct {

}

func (*Estructuras) Mifuncion()string {
	return "Hola desde mi funcion"
}

func (*Estructuras) calculo(n1 int, n2 int) int { 
	return n1 + n2
}

/*
//estructuras
func main() {
	fmt.Println("Hola Mundo desde Go en la Clase 3")
	estructuras := Persona{
		id:     1,
		nombre: "Rids",
		correo:  "mail.com",
		edad:   25,
	}

	fmt.Println(estructuras)
	fmt.Printf("%+v \n", estructuras)
	fmt.Println(reflect.TypeOf(estructuras))
	//2do forma
	p:= new (Persona)
	fmt.Println(reflect.TypeOf(p))
	p.nombre="Ridss"
	p.edad=26
	p.correo="@mail.com"
	p.id=2
	fmt.Println(p)
	fmt.Printf("%+v \n", p)

	//Estructuras anidadas
	categoria:= categoria{
		id:1,
		nombre:"Tecnologia",
		slug:"tecno",
	}
	
	producto:= producto{
		id:1,
		nombre:"Monitor",
		slug:"monitor-24p",
		precio:300,
		categoria:categoria,
	}
	fmt.Printf("%+v \n", producto)

}

//Persona -> publica
//persona -> privada
type Persona struct {
	id    int
	nombre string
	correo  string
	edad   int
}

type categoria struct {
	id   int
	nombre string
	slug   string
}

type producto struct {
	id         int
	nombre     string
	slug	   string
	precio	 int
	categoria categoria
}
*/