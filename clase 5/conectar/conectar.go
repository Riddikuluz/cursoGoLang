package conectar

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB

func Conectar() {

	errorVariables := godotenv.Load();
	if errorVariables != nil {
		panic( errorVariables)
	}

	// Crear DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Abrir conexión
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al abrir conexión: %v", err)
	}

	// Probar conexión
	if err := conn.Ping(); err != nil {
		log.Fatalf("Error al hacer ping: %v", err)
	}

	Db = conn
	fmt.Println("Conexión a DB OK")
}

func CerrarConexion() {
	if Db != nil {
		Db.Close()
	}
}
