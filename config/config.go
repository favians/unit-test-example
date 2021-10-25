package config

import (
	"fmt"
	"mvc/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dbName := "alta13"

	if os.Getenv("ENV") == "test" {
		dbName = dbName + "_test"
	}

	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "yourpassword",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     dbName,
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.Users{})
}
