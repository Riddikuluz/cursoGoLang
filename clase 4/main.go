package main

import (
	"c4/moduloEjemplo"
	"fmt"
)

//modulo personalizado
func main() {
	saludo := moduloEjemplo.Saludar("Carlos")
	fmt.Println(saludo)

	despedida := moduloEjemplo.Despedir("Carlos")
	fmt.Println(despedida)
}



/*
//logs
func main() {

	//fatal detiene la ejecucion del programa
	//error no

	//err := errors.New("este es un error")	
	//log.Fatal(err)
	//log.Fatal("hola soy un error")
	//fmt.Println("error")
	//log.Println("esto es un log")

	//log.Panicln("panic at the disco")

	//guarar logs en un archivo
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("No se pudo abrir el archivo de log:", err)
	}
	defer file.Close()

	log.SetOutput(file)

	log.Println("Este es un mensaje de log")
	log.Println("Otro mensaje de log")

	//simular error
	err = errors.New("este es un error simulado")
	if err != nil {
		log.Println("Error ocurrido:", err)
	}
}		

//os
func main() {
	nombre := flag.String("nombre", "", "nombre de la persona")
	edad := flag.Int("edad", 18, "edad de la persona")
	flag.Parse()
	//espacio en memoria
	fmt.Println("Nombre:", nombre)
	fmt.Println("Edad:", edad)
	//punteros
	fmt.Println("Nombre:", *nombre)
	fmt.Println("Edad:", *edad)

}


//math/rand
func main(){
	//aleatorio entre 0 y 100
	num := rand.Intn(101)
	fmt.Println("Número aleatorio entre 0 y 100:", num)

	min := 50
	max := 150
	rand.Seed(time.Now().UnixNano())
	aleatorio2 := rand.Intn(max - min + 1) + min
	fmt.Printf("Número aleatorio entre %d y %d: %d\n", min, max, aleatorio2)

	//password
	longitud := 128
	password := generarPassword(longitud)
	fmt.Println("Contraseña generada:", password)	
}

func generarPassword(longitud int) string {
	const caracteres = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
	password := make([]byte, longitud)
	rand.Seed(time.Now().UnixNano())
	for i := range password {
		password[i] = caracteres[rand.Intn(len(caracteres))]
	}
	return string(password)
}

//strings
func main() {
	cadena := "Hola, ¿cómo estás?"
	fmt.Println("Cadena original:", cadena)

	// Convertir a mayúsculas
	mayusculas := strings.ToUpper(cadena)
	fmt.Println("Mayúsculas:", mayusculas)

	// Convertir a minúsculas
	minusculas := strings.ToLower(cadena)
	fmt.Println("Minúsculas:", minusculas)

	//split
	palabras := strings.Split(cadena, " ")
	fmt.Println("Palabras:", palabras)

	//buscar subcadena
	subcadena := "cómo"
	posicion := strings.Index(cadena, subcadena)
	if posicion != -1 {
		fmt.Printf("La subcadena '%s' se encuentra en la posición %d\n", subcadena, posicion)
	} else {
		fmt.Printf("La subcadena '%s' no se encontró\n", subcadena)
	}

	//repetidas	
	repetida := "Hola "
	repeticiones := 2
	resultado := strings.Repeat(repetida, repeticiones)
	fmt.Printf("Cadena repetida %d veces: %s\n", repeticiones, resultado)

	//remplazar palabra
	original := "estás"
	nuevo := "te encuentras"
	cadenaModificada := strings.Replace(cadena, original, nuevo, 1)
	fmt.Println("Cadena modificada:", cadenaModificada)

}

//modulo time
func main() {
	// Obtener la hora actual
	now := time.Now()
	fmt.Println("Hora actual (sin formato):", now)

	fmt.Println(now.Year())	
	fmt.Println(now.Month())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())	
	fmt.Println(now.Weekday())	
	fmt.Println(now.Hour())	
	fmt.Println(now.Minute())	
	fmt.Println(now.Second())

	//sumar dias
	future := now.AddDate(0, 0, 10) // Sumar 10 días
	fmt.Println("Fecha dentro de 10 días:", future)

	//restar dias
	past := now.AddDate(0, 0, -10) // Restar 10 días
	fmt.Println("Fecha hace 10 días:", past)

	formatDates(now)
}	

func formatDates(now time.Time) {
	// Formatear la fecha en diferentes formatos
	fmt.Println("Formato RFC3339:", now.Format(time.RFC3339))
	fmt.Println("Formato ANSIC:", now.Format(time.ANSIC))
	fmt.Println("Formato UnixDate:", now.Format(time.UnixDate))
	fmt.Println("Formato personalizado:", now.Format("02-01-2006 15:04:05"))

}
*/