package configs

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB{
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", 
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

	migrateDDL(db)
	initialDataDML(db)
	return db
}

func migrateDDL(db *gorm.DB) {
	db.AutoMigrate(&entities.Category{})
	db.AutoMigrate(&entities.Item{})
	db.AutoMigrate(&entities.User{})
}

func initialDataDML(db *gorm.DB) {
	query, err := ioutil.ReadFile("db/initial_data.sql")

	if err != nil {
		panic(err)
	}
	
	var count int64
	db.Model(&entities.Category{}).Count(&count)
	
	if count > 0 {
		return
	}

	statements := strings.Split(string(query), ";")
	
	for _, statement := range statements {
		if statement == "" {
			continue
		}
		
		if err := db.Exec(statement).Error; err != nil {
			panic(err)
		}
	}

}