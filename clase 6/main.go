package main

import (
	"c6/rutas"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//go get -u github.com/gorilla/mux
//go get github.com/joho/godotenv
//go install github.com/gravityblast/fresh@latest
//go run github.com/gravityblast/fresh@latest

func main() {
	mux:= mux.NewRouter()
	//rutas
	mux.HandleFunc("/", rutas.Home)
	mux.HandleFunc("/nosotros", rutas.Nosotros)
	mux.HandleFunc("/parametros/{id:.*}/{slug:.*}", rutas.Parametros)
	mux.HandleFunc("/querystring", rutas.QueryString)
	mux.HandleFunc("/estructuras", rutas.Estructuras)

	//archivos estaticos mux
	s:= http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	mux.PathPrefix("/public/").Handler(s)

	//error 404
	mux.NotFoundHandler = mux.NewRoute().HandlerFunc(rutas.Pagina404).GetHandler()

	//cargar variables de entorno
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}
	//servidor
	server:= &http.Server{	
		Addr: os.Getenv("ADDR")+os.Getenv("PORT"),
		Handler: mux,
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15*time.Second,
	}
	fmt.Println("corriendo en http://"+os.Getenv("ADDR")+os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())
}

/*
func main() {
	//mux:= http.NewServeMux()

	http.HandleFunc("/", func(response http.ResponseWriter, Request *http.Request) {
		fmt.Fprintln(response, "Hola desde mi primer servidor en Go")
	})

	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}
*/