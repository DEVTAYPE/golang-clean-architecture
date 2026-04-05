package database

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(connectionString string) error {
	var err error

	// Realizamos la conexión a la base de datos utilizando el connection string proporcionado
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return errors.New("Error al conectar a la base de datos: " + err.Error())
	}

	// Hacemos ping para verificar que la conexión es válida sobre el DB pool
	err = DB.Ping()
	if err != nil {
		return errors.New("Error al verificar la conexión a la base de datos: " + err.Error())
	}

	// establecemos el maximo de conexiones abiertas y en espera
	DB.SetMaxOpenConns(25) // Conexiones abiertas máximas
	DB.SetMaxIdleConns(25) // Conexiones en espera

	return nil
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}

	return nil
}
