package configs

import (
	"fmt"
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
	initialDataCategoryDML(db)
	return db
}

func migrateDDL(db *gorm.DB) {
	db.AutoMigrate(&entities.Category{})
	db.AutoMigrate(&entities.Item{})
	db.AutoMigrate(&entities.User{})
}

func initialDataCategoryDML(db *gorm.DB) {

	var count int64
	db.Model(&entities.Category{}).Count(&count)
	
	if count > 0 {
		return
	}
	
	categories :=  []string{"Laptop", "Computer", "Accessories", "Mobile", "Tablet", "Camera", "TV"}

	for _, category := range categories {
		db.Create(&entities.Category{
			Name: strings.ToLower(category),
		})
	}

}