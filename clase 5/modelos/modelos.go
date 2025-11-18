package modelos

type Cliente struct {
	Id    int
	Nombre string
	correo  string
	Telefono string
	Fecha_registro string
}

type Clientes []Cliente