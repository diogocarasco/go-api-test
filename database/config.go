package database

// getDsn retrieve the DSN
func getDsn() string {

	//dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_ROOT_HOST"), os.Getenv("MYSQL_DATABASE"))

	dsn := "postgresql://postgres:postgres@localhost:5432"

	return dsn

}
