package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/diogocarasco/go-api-test/models"
)

var DB *gorm.DB
var Mock sqlmock.Sqlmock

// InitializeDB create a new connection to the database
func InitializeDB(mode string) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	if mode == "mock" {

		var db *sql.DB
		var err error

		db, Mock, err = sqlmock.New()
		if err != nil {
			fmt.Printf("an error '%s' was not expected when opening a stub database connection", err)
		}

		dialector := postgres.New(postgres.Config{
			DSN:                  "sqlmock_db_0",
			DriverName:           "postgres",
			Conn:                 db,
			PreferSimpleProtocol: true,
		})
		DB, err = gorm.Open(dialector, &gorm.Config{})
		if err != nil {
			fmt.Printf("an error '%s' was occurred", err)
		}

	} else {

		for i := 1; i <= 3; i++ {
			//DSN: "host=localhost user=postgres password=postgres database=spfeiras port=5432 sslmode=disable",
			db, err := gorm.Open(postgres.New(postgres.Config{
				DSN: getDsn(),
			}), &gorm.Config{Logger: newLogger})

			if err != nil {
				fmt.Printf("[Try %d] Database connection error: %s \n", i, err)
				time.Sleep(5 * time.Second)
			} else {
				fmt.Printf("[Try %d] Database connected!\n", i)
				db.AutoMigrate(&models.Feiras{})
				DB = db
				break
			}

		}

	}

}
