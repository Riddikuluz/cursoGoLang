package main

import (
	"fmt"
	"reflect"
)

//mapa
func main() {
	paises := make(map[string]int)
	paises["Chile"] = 128932753
	paises["colombia"] = 50882891
	paises["argentina"] = 45195777
	paises["peru"] = 32971846
	

	fmt.Println("mapa paises:", paises)
	fmt.Println("poblacion de mexico:", paises["Chile"])
	fmt.Println("longitud del mapa paises:", len(paises))
	fmt.Printf("tipo de dato del mapa paises: %s \n", reflect.TypeOf(paises))

	paises2 := map[int]string{
		1: "Chile",
		2: "colombia",
		3: "argentina",
		4: "peru",
	}
	
	fmt.Println("mapa paises2:", paises2)
	fmt.Println("pais con clave 2:", paises2[2])
	fmt.Println("longitud del mapa paises2:", len(paises2))
	fmt.Printf("tipo de dato del mapa paises2: %s \n", reflect.TypeOf(paises2))

	// revisar si una clave existe en un mapa
	pais, existe := paises2[11]
	if existe {
		fmt.Println("la poblacion de bolivia es:", pais)
	} else {
		fmt.Println("la clave bolivia no existe en el mapa paises")
	}

	//eliminar un elemento de un mapa
	delete(paises, "peru")
	fmt.Println("mapa paises despues de eliminar un elemento:", paises)
	fmt.Println("longitud del mapa paises despues de eliminar un elemento:", len(paises))

	//recorrer un mapa con for range
	for id, valor := range paises2 {
		fmt.Printf("clave: %d, valor: %s \n", id, valor)
	}

	//json
	respuesta := map[string]string{
		"nombre": "rids",
		"estado": "ok",
	}
	fmt.Println("respuesta json:", respuesta)
	fmt.Println("estado: ", respuesta["estado"])

}

/*	
//areglos y slices
func main() {
	//arreglos
	var paises [4]string
	paises[0] = "Chile"
	paises[1] = "colombia"
	paises[2] = "peru"
	paises[3] = "argentina"
	fmt.Println("arreglo paises:", paises)
	fmt.Println("primer elemento del arreglo paises:", paises[0])
	fmt.Println("longitud del arreglo paises:", len(paises))
	fmt.Printf("tipo de dato del arreglo paises: %s \n", reflect.TypeOf(paises))

	//slices
	var paises2 =  make([]string, 5)
	paises2[0] = "mexico"
	paises2[1] = "bolivia"
	paises2[2] = "ecuador"
	paises2[3] = "venezuela"
	paises2[4] = "uruguay"
	fmt.Println("slice paises2:", paises2)
	fmt.Println("primer elemento del slice paises2:", paises2[0])
	fmt.Println("longitud del slice paises2:", len(paises2))
	fmt.Printf("tipo de dato del slice paises2: %s \n", reflect.TypeOf(paises2))
	
	// agregar elementos a un slice
	paises2 = append(paises2, "paraguay")
	fmt.Println("slice paises2 despues de agregar un elemento:", paises2)
	fmt.Println("longitud del slice paises2 despues de agregar un elemento:", len(paises2))

	//eliminar un elemento de un slice
	paises2 = append(paises2[:2], paises2[2+1:]...) //elimina el elemento en la posicion 2
	fmt.Println("slice paises2 despues de eliminar un elemento:", paises2)
	fmt.Println("longitud del slice paises2 despues de eliminar un elemento:", len(paises2))

}


//ciclos
func main() {
	i := 0
	for i < 5 {
		fmt.Println("el valor de i es:", i)
		i++
	}

	for i2:= 1; i2 <= 5; i2++ {
		if i2 == 3 {
			//break
			continue
		}
		fmt.Println("el valor de i2 es:", i2)
		
	}

}

func main() {
	edad := 30
	if edad >= 18 {
		fmt.Println("es mayor de edad")
	} else {
		fmt.Println("es menor de edad")
	}

	color := "azul"
	if color == "rojo" {
		fmt.Println("el color es rojo")
	} else if color == "verde" {
		fmt.Println("el color es verde")
	} else if color == "azul" {
		fmt.Println("el color es azul")
	} else {
		fmt.Println("el color no es rojo, verde o azul")
	}
	//operador logico AND &&
	if color == "rojo" && edad >= 18 {
		fmt.Println("el color es rojo y es mayor de edad")
	} else {
		fmt.Println("el color no es rojo o no es mayor de edad")
	}
	//declaracion de variable dentro del if
	if variable := 1; variable > 5 {
		fmt.Println("la variable es mayor que 5")
	} else {
		fmt.Println("la variable es menor o igual que 5")
	}

	//switch case
	switch color {
	case "rojo":
		fmt.Println("el color es rojo")
	case "verde":
		fmt.Println("el color es verde")
	case "azul":
		fmt.Println("el color es azul")
	default:
		fmt.Println("el color no es rojo, verde o azul")
	}

}


func main() {
	color := "rojo"
	fmt.Println(color,&color)
	estado := true
	fmt.Println(estado,&estado)
	}			


//reflect y tipeof
//import reflect
func main() {
	var string1  = 111111111111111111
	fmt.Println(string1)
	fmt.Printf("el tipo de la variable es: %s \n", reflect.TypeOf(string1))
}

func main() {
	var string1 string = "hola mundo"
	fmt.Println(string1)
	textoGrande := `esto es un texto`
	fmt.Println(textoGrande)
	var estado bool = true	
	fmt.Println(estado)
	var flotante32 float32 = 12.45
	fmt.Println(flotante32)
	var flotante64 float64 = 123456789.123456789
	fmt.Println(flotante64)
	var entero int = 123456
	fmt.Println(entero)
	var entero8 int8 = 127
	fmt.Println(entero8)
	var entero16 int16 = 32767
	fmt.Println(entero16)
	var entero32 int32 = 2147483647
	fmt.Println(entero32)
	var entero64 int64 = 9223372036854775807
	fmt.Println(entero64)
	var untero uint = 123456
	fmt.Println(untero)
	}	

	
//constantes
const miConstante = "hola constante"

// varibbles
func main() {
	var nombre string = "rids"
	fmt.Println(nombre)
	nombre2 := "cesar"
	fmt.Println(nombre2)
	fmt.Printf("el valor es %s", miConstante)

}
*/