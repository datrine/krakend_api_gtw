package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func NotReady() {
	byt, err := os.ReadFile(".env")
	if err != nil {
		panic(err.Error())
	}
	str := string(byt)
	strA := strings.Split(str, "\n")
	for _, s := range strA {
		st := strings.TrimSpace(s)
		if strings.Index(st, "#") == 0 {
			continue
		}
		subA := strings.Split(st, "=")
		if len(subA) <= 1 {
			panic(fmt.Errorf("bad format for .env"))
		}
		key := subA[0]
		val := strings.Join(subA[1:], "=")
		if strings.Index(val, "\"") == 0 && strings.Index(val, "\"") == len(val)-1 {
			val = val[1 : len(val)-1]
		}
		err = os.Setenv(key, val)
		if err != nil {
			panic(err.Error())
		}
		println(key, val)
	}
}

func GetPort() int {
	var str string = os.Getenv("PORT")
	if str == "" {
		return 3000
	}
	port, err := strconv.Atoi(str)
	if err != nil {
		panic(err.Error())
	}
	return port
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
