package rutas

import (
	"c6/utilidades"
	"c6/validaciones"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)
	
func Formularios_get(response http.ResponseWriter, request *http.Request) {
	template:= template.Must(template.ParseFiles("templates/formularios/formularios.html", utilidades.Frontend))
	css_sesion, mensaje_sesion:= utilidades.RetornarMensajesFlash(response, request)
	data:= map[string]interface{}{
		"css": css_sesion,
		"mensaje": mensaje_sesion,
	}
	template.Execute(response, data)
}

func Formularios_post(response http.ResponseWriter, request *http.Request) {
	mensaje:= ""
	if len(request.FormValue("nombre")) == 0 {
		mensaje = mensaje + "El campo nombre es obligatorio. "
	} else if len(request.FormValue("correo")) == 0 {
		mensaje = mensaje + "El campo correo es obligatorio. "
	} else if len(request.FormValue("telefono")) == 0 {
		mensaje = mensaje + "El campo telefono es obligatorio. "
	} else if len(request.FormValue("contrasena")) == 0 {
		mensaje = mensaje + "El campo contrasena es obligatorio. "
	}
	if validaciones.Regex_correo.FindStringSubmatch(request.FormValue("correo"))== nil {
		mensaje = mensaje + "El campo correo no es valido. "
	}
	if validaciones.ValidarPassword(request.FormValue("contrasena")) == false {
		mensaje = mensaje + "El campo contrasena no es valido. Debe tener al menos 8 caracteres, una mayuscula, una minuscula, un numero y un caracter especial. "
	}	

	if mensaje != "" {
		//fmt.Fprintln(response, mensaje)
		//return
		utilidades.CrearMensajeFlash(response, request,"danger", mensaje)
		http.Redirect(response, request, "/formularios", http.StatusSeeOther)
	}
		fmt.Fprintln(response, "nombre:", request.FormValue("nombre"))
		fmt.Fprintln(response, "correo:", request.FormValue("correo"))
		fmt.Fprintln(response, "telefono:", request.FormValue("telefono"))
		fmt.Fprintln(response, "contrasena:", request.FormValue("contrasena"))
	
}

func Formularios_upload(response http.ResponseWriter, request *http.Request) {
	template := template.Must(template.ParseFiles("templates/formularios/upload.html", utilidades.Frontend))
	css_sesion, css_mensaje := utilidades.RetornarMensajesFlash(response, request)
	data := map[string]string{

		"css":     css_sesion,
		"mensaje": css_mensaje,
	}
	template.Execute(response, data)

}

func Formularios_upload_post(response http.ResponseWriter, request *http.Request) {
	file, handler, err := request.FormFile("foto")
	if err != nil {
		utilidades.CrearMensajeFlash(response, request, "danger", "Ocurrió un error inesperado 1")
	}
	var extension = strings.Split(handler.Filename, ".")[1]
	time := strings.Split(time.Now().String(), " ")
	foto := string(time[4][6:14]) + "." + extension
	var archivo string = "public/upload/fotos/" + foto
	f, errCopy := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0777)
	if errCopy != nil {
		utilidades.CrearMensajeFlash(response, request, "danger", "Ocurrió un error inesperado 2")
	}
	_, errCopiar := io.Copy(f, file)
	if errCopiar != nil {
		utilidades.CrearMensajeFlash(response, request, "danger", "Ocurrió un error inesperado 3")
	}
	//acá lo guardarías en la bd

	//redireccionamos
	utilidades.CrearMensajeFlash(response, request, "success", "Se subió el archivo "+foto+" exitosamente")
	http.Redirect(response, request, "/formularios/upload", http.StatusSeeOther)

}


/*
func Formularios_post(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "nombre:", request.FormValue("nombre"))
	fmt.Fprintln(response, "correo:", request.FormValue("correo"))
	fmt.Fprintln(response, "telefono:", request.FormValue("telefono"))
	fmt.Fprintln(response, "contrasena:", request.FormValue("contrasena"))
}
*/