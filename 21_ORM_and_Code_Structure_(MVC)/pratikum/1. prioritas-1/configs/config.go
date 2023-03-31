package configs

import (
	"fmt"

	cons "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/constants"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cons.DB_USERNAME,
		cons.DB_PASSWORD,
		cons.DB_HOST,
		cons.DB_PORT,
		cons.DB_NAME,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	initialMigration()
}

func initialMigration() {
  DB.AutoMigrate(&models.User{})
}

