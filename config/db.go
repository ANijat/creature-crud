package config

import (
	"fmt"
	"os"
)

var (
	DBUser     = os.Getenv("DBUser")
	DBPassword = os.Getenv("DBPassword")
	DBName     = os.Getenv("DBName")
	DBHost     = os.Getenv("DBHost")
	DBPort     = os.Getenv("DBPort")
	DBtype     = os.Getenv("DBtype")
)

func GetDBType() string {
	return DBtype
}

func GetDBConnectionString() string {
	dataBase := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser,
		DBPassword,
		DBHost,
		DBPort,
		DBName)

	return dataBase
}
