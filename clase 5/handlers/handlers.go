package handlers

import (
	"bufio"
	"c5/conectar"
	"fmt"
	"os"
)

func Listar() {
	conectar.Conectar()
	defer conectar.CerrarConexion()
	sql := "SELECT id, nombre, correo, telefono, fecha_registro FROM cliente"
	rows, err := conectar.Db.Query(sql)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:",  err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nombre, correo, telefono, fecha_registro string
		if err := rows.Scan(&id, &nombre, &correo, &telefono, &fecha_registro); err != nil {
			fmt.Println("Error al escanear fila:", err)
			return
		}
		//fmt.Println(rows)
		fmt.Printf("ID: %d, Nombre: %s, Correo: %s, Teléfono: %s, Fecha registro: %s \n", id, nombre, correo, telefono, fecha_registro)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error durante la iteración de filas:", err)
	}	

	
}

func ListarId(id int) {
	conectar.Conectar()
	defer conectar.CerrarConexion()
	sql := "SELECT id, nombre, correo, telefono, fecha_registro FROM cliente WHERE id = ?"
	row := conectar.Db.QueryRow(sql, id)

	var nombre, correo, telefono, fecha_registro string
	if err := row.Scan(&id, &nombre, &correo, &telefono, &fecha_registro); err != nil {
		fmt.Println("Error al escanear fila:", err)
		return
	}
	fmt.Printf("ID: %d, Nombre: %s, Correo: %s, Teléfono: %s, Fecha registro: %s \n", id, nombre, correo, telefono, fecha_registro)

	
}

func Insertar(nombre string, correo string, telefono string) {
	conectar.Conectar()
	defer conectar.CerrarConexion()
	sql := "INSERT INTO cliente (nombre, correo, telefono) VALUES (?, ?, ?)"

	result, err := conectar.Db.Exec(sql, nombre, correo, telefono)
	if err != nil {
        fmt.Println("Error al eliminar:", err)
        return
    }
	
	fmt.Println(result, err)
	
}

func Actualizar(id int, nombre string, correo string, telefono string) {
	conectar.Conectar()
	defer conectar.CerrarConexion()
	sql := "UPDATE cliente SET nombre = ?, correo = ?, telefono = ? WHERE id = ?"

	result, err := conectar.Db.Exec(sql, nombre, correo, telefono, id)

	if err != nil {
        fmt.Println("Error al eliminar:", err)
        return
    }
	
	fmt.Println(result, err)
	
}

func Eliminar(id int) {
	conectar.Conectar()
	defer conectar.CerrarConexion()
	sql := "DELETE FROM cliente WHERE id = ?"

	result, err := conectar.Db.Exec(sql, id)

	if err != nil {
        fmt.Println("Error al eliminar:", err)
        return
    }

	fmt.Println(result, err)
	
}	

func Ejecutar () {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("1 - listar")
	fmt.Println("2 - listar por id")
	fmt.Println("3 - insertar")
	fmt.Println("4 - actualizar")
	fmt.Println("5 - eliminar")
	fmt.Println("Ingrese un comando SQL:")
	if scanner.Scan() {
		for {
			if scanner.Text() == "1" {
				Listar()
				break
			} else if scanner.Text() == "2" {
				fmt.Println("Ingrese el ID:")
				if scanner.Scan() {
					var id int
					fmt.Sscanf(scanner.Text(), "%d", &id)
					ListarId(id)
					break
				}
			} else if scanner.Text() == "3" {
				fmt.Println("Ingrese nombre, correo y teléfono separados por espacios:")
				
				var nombre, correo, telefono string
				_, err := fmt.Scanln(&nombre, &correo, &telefono)
				if err != nil {
					fmt.Println("Error al leer la entrada:", err)
					break
				}
				Insertar(nombre, correo, telefono)
				break
				
			} else if scanner.Text() == "4" {
				fmt.Println("Ingrese id, nombre, correo y teléfono separados por espacios:")
				var id int
				var nombre, correo, telefono string
				_, err := fmt.Scanln(&id, &nombre, &correo, &telefono)
				if err != nil {
					fmt.Println("Error al leer la entrada:", err)
					break
				}
				Actualizar(id, nombre, correo, telefono)
				break

			} else if scanner.Text() == "5" {
				fmt.Println("Ingrese el ID:")
				if scanner.Scan() {	
					var id int
					fmt.Sscanf(scanner.Text(), "%d", &id)
					Eliminar(id)
					break
				}
			}
		
		}
	}

}