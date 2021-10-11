package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

// Initiate auto migrate the database
func Initiate() {
	dsn := "host=localhost user=postgres password=12qwaszx dbname=postgres port=4101 sslmode=disable TimeZone=Asia/Jakarta"
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
