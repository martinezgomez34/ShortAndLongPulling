package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

type ConnMySQL struct {
	DB  *sql.DB
	Err error
}

func GetDBPool() *ConnMySQL {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbSchema := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbSchema)

	
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
		return nil
	}

	db.SetMaxOpenConns(10)
	
	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Error al verificar la conexión: %v", err)
		return nil
	}

	fmt.Println("Conexión a la base de datos establecida correctamente")
	return &ConnMySQL{DB: db}
}

func (conn *ConnMySQL) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	if conn.DB == nil {
		return nil, fmt.Errorf("la conexión a la base de datos no está establecida")
	}
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("Error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta preparada: %w", err)
	}

	return result, nil
}

func (conn *ConnMySQL) FetchRows(query string, values ...interface{}) (*sql.Rows, error) {
	if conn.DB == nil {
		return nil, fmt.Errorf("la conexión a la base de datos no está establecida")
	}
	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %w", err)
	}

	return rows, nil
}
