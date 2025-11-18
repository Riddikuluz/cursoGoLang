package rutas

import (
	"c6/utilidades"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) {
	template:= template.Must(template.ParseFiles("templates/ejemplos/home.html", utilidades.Frontend))
	template.Execute(response, nil)

}

func Nosotros(response http.ResponseWriter, request *http.Request) {
	template:= template.Must(template.ParseFiles("templates/ejemplos/nosotros.html", utilidades.Frontend))
	template.Execute(response, nil)

}
func Parametros(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	data:= map[string]string{
		"id": vars["id"],
		"slug": vars["slug"],
	}
	template:= template.Must(template.ParseFiles("templates/ejemplos/parametros.html", utilidades.Frontend))
	template.Execute(response, data)

}
func QueryString(response http.ResponseWriter, request *http.Request) {
	data:= map[string]string{
		"id": request.URL.Query().Get("id"),
		"slug": request.URL.Query().Get("slug"),
	}
	template:= template.Must(template.ParseFiles("templates/ejemplos/queryString.html", utilidades.Frontend))
	template.Execute(response, data)

}

type Habilidad struct {
	Nombre string
}

type Datos struct {
	Nombre     string
	Edad       int
	Perfil     int
	Habilidades []Habilidad
}

func Estructuras(response http.ResponseWriter, request *http.Request) {
	habilidad1:= Habilidad{"Programación"}
	habilidad2:= Habilidad{"Diseño web"}
	habilidad3:= Habilidad{"Base de datos"}

	datos:= Datos{
		Nombre: "Rids",
		Edad: 25,
		Perfil: 1,
		Habilidades: []Habilidad{habilidad1, habilidad2, habilidad3},
	}
	template:= template.Must(template.ParseFiles("templates/ejemplos/estructuras.html", utilidades.Frontend))
	template.Execute(response, datos)
}

func Pagina404(response http.ResponseWriter, request *http.Request) {
	template:= template.Must(template.ParseFiles("templates/ejemplos/pagina404.html", utilidades.Frontend))
	template.Execute(response, nil)

}

/*
func Home(response http.ResponseWriter, request *http.Request) {
	template, err:= template.ParseFiles("templates/ejemplos/home.html", "templates/layout/fronend.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}

func Nosotros(response http.ResponseWriter, request *http.Request) {
	template, err:= template.ParseFiles("templates/ejemplos/nosotros.html", "templates/layout/fronend.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, nil)
	}
}

func Parametros(response http.ResponseWriter, request *http.Request) {
	template, err:= template.ParseFiles("templates/ejemplos/parametros.html", "templates/layout/fronend.html")
	vars := mux.Vars(request)
	data:= map[string]string{
		"id": vars["id"],
		"slug": vars["slug"],
	}
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, data)
	}
}

func QueryString(response http.ResponseWriter, request *http.Request) {
	template, err:= template.ParseFiles("templates/ejemplos/queryString.html", "templates/layout/fronend.html")
	data:= map[string]string{
		"id": request.URL.Query().Get("id"),
		"slug": request.URL.Query().Get("slug"),
	}
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, data)
	}
}
type Habilidad struct {
	Nombre string
}

type Datos struct {
	Nombre     string
	Edad       int
	Perfil     int
	Habilidades []Habilidad
}

func Estructuras(response http.ResponseWriter, request *http.Request) {
	template, err:= template.ParseFiles("templates/ejemplos/estructuras.html", "templates/layout/fronend.html")
	habilidad1:= Habilidad{"Programación"}
	habilidad2:= Habilidad{"Diseño web"}
	habilidad3:= Habilidad{"Base de datos"}

	datos:= Datos{
		Nombre: "Rids",
		Edad: 25,
		Perfil: 1,
		Habilidades: []Habilidad{habilidad1, habilidad2, habilidad3},
	}

	if err != nil {
		panic(err)
	} else {
		template.Execute(response, datos)
	}
}
*/

/*
func Estructuras(response http.ResponseWriter, request *http.Request) {
	template, err:= template.ParseFiles("templates/ejemplos/estructuras.html")
	if err != nil {
		panic(err)
	} else {
		template.Execute(response, Datos{Nombre: "Rids", Edad: 25, Perfil: 1})
	}
}

func Home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hola desde mi primer servidor en Go usando Gorilla Mux")
}

func Nosotros(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hola desde la sección de Nosotros con fresh")
}

func Parametros(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprintln(response, "ID = "+vars["id"]+" | Slug = "+vars["slug"])
}

func QueryString(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, request.URL)
	fmt.Fprintln(response, request.URL.RawQuery)
	fmt.Fprintln(response, request.URL.Query())
	fmt.Fprintln(response, request.URL.Query().Get("id"))
	fmt.Fprintln(response, request.URL.Query().Get("slug"))
}
*/