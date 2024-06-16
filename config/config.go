package config

import (
	"fmt"
	"os"
)

func init() {

}

func GetLocalDBName() string {
	str := os.Getenv("LOCAL_DB_NAME")
	if str == "" {
		return "local.db"
	}
	return str
}

func GetTursoDatabaseURL() string {
	str := os.Getenv("TURSO_DATABASE_URL")
	if str == "" {
		panic(fmt.Errorf("TURSO_DATABASE_URL must be set"))
	}
	return str
}

func GetTursoAuthToken() string {
	str := os.Getenv("TURSO_AUTH_TOKEN")
	if str == "" {
		panic(fmt.Errorf("TURSO_AUTH_TOKEN must be set"))
	}
	return str
}
