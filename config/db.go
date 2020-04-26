package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database instance
var DB *gorm.DB

// Database settings
const (
	host     = "localhost"
	port     = 5432 // Default port
	user     = "postgres"
	password = "password"
	dbname   = "golpo_fiber"
)

// ConnectDB function
func ConnectDB() error {
	var err error
	DB, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}
	//if err = DB.Ping(); err != nil {
	//	return err
	//}
	return nil
}
