package configs

import (
	"fmt"
	"log"
	"os"

	models_mysql "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init(){
	initDB()
}

func initDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	initialMigration()
}

func initialMigration() {
  DB.AutoMigrate(&models_mysql.User{})
	DB.AutoMigrate(&models_mysql.Book{})
	DB.AutoMigrate(&models_mysql.Blog{})
}
