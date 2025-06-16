package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DbContext *sql.DB

//========================= ENV =========================

func InitEnv() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		fmt.Println(Red, "error .env : ", errEnv, Reset)
		return
	}
	if GetEnvWithDefault("DB_NAME", "") == "" || GetEnvWithDefault("DB_USER", "") == "" || GetEnvWithDefault("DB_PWD", "") == "" || GetEnvWithDefault("DB_PORT", "") == "" || GetEnvWithDefault("DB_HOST", "") == "" {
		log.Fatalf(Red, "Error with GetEnWithDefault, impossible launch server", Reset)
	}

	DB_Name, CheckDB_Name := os.LookupEnv("DB_Name")
	if !CheckDB_Name {
		fmt.Println(Red, "error CheckDB_Name : %s", DB_Name, Reset)
		return
	}
	DB_Port, CheckDB_Port := os.LookupEnv("DB_Port")
	if !CheckDB_Port {
		fmt.Println(Red, "error CheckDB_Port : %s", DB_Port, Reset)
		return
	}

}

//========================= DB =========================

//------------------ init ------------------

func InitDB() {
	user := GetEnvWithDefault("DB_USER", "")
	pwd := GetEnvWithDefault("DB_PWD", "")
	host := GetEnvWithDefault("DB_HOST", "localhost")
	port := GetEnvWithDefault("DB_PORT", "4000")
	name := GetEnvWithDefault("DB_NAME", "")

	if user == "" || name == "" {
		log.Fatalf(Red, "Error with user or name, impossible launch DB and server", Reset)
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pwd, host, port, name)
	var err error
	DbContext, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf(Red, "Erreur d'ouverture de la connexion : %v", err, Reset)
	}

	err = DbContext.Ping()
	if err != nil {
		DbContext.Close()
		log.Fatalf(Red, "Fail ping to MySQL, error : %v", err, Reset)
	}
}



func GetEnvWithDefault(key string, defaultValue string) string {
	envVar, envErr := os.LookupEnv(key)
	if !envErr {
		return defaultValue
	}
	return envVar
}
