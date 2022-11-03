package database

import (
	"fmt"
	"log"
	"os"
	"time"

	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/diogocarasco/go-api-test/models"
)

var DB *gorm.DB

// InitializeDB create a new connection to the database
func InitializeDB() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	for i := 1; i <= 3; i++ {

		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN: "host=localhost user=postgres password=postgres database=spfeiras port=5432 sslmode=disable",
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
