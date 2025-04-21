package conectar

import (
	"database/sql"
	"os"

	"github.com/KevinCas1401/ejerciciomysql"
)

var DB *sql.DB

func conectar() {
	errorVariable := ejerciciomysql.Load()

	if errorVariable != nil {
		panic(errorVariable)

	}

	conection, err := sql.Open("mysyql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}
	DB = conection

}

func terminarconexion() {
	DB.Close()
}
