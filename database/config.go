package database

import (
	"fmt"
	"os"
)

// getDsn retrieve the DSN
func getDsn() string {

	//dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_ROOT_HOST"), os.Getenv("MYSQL_DATABASE"))

	//dsn := "postgresql://postgres:postgres@localhost:5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"))

	//"host=localhost user=postgres password=postgres database=spfeiras port=5432 sslmode=disable"
	return dsn

}
