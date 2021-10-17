package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

// Initiate auto migrate the database
func Initiate() {
	// load from env
	postgresHost := os.Getenv("POSTGRESQL_HOST")
	postgresUser := os.Getenv("POSTGRESQL_USER")
	postgresPassword := os.Getenv("POSTGRESQL_PASSWORD")
	postgresDatabase := os.Getenv("POSTGRESQL_DATABASE")
	postgresPort := os.Getenv("POSTGRESQL_PORT")
	postgresSslMode := os.Getenv("POSTGRESQL_SSL_MODE")
	postgresTimezone := os.Getenv("POSTGRESQL_TIMEZONE")
	dsn := "host=" + postgresHost + " user=" + postgresUser + " password=" + postgresPassword +
		" dbname=" + postgresDatabase + " port=" + postgresPort + " sslmode=" + postgresSslMode +
		" TimeZone=" + postgresTimezone

	var err error
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Printf("%#v\n", DBConn)

	// add models here
	fmt.Println("Initiate Database")
	DBConn.AutoMigrate(&User{})
}
