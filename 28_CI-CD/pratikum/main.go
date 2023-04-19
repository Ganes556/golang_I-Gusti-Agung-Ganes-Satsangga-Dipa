package main

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/config"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/route"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
	
	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	route.NewRoute(e, db)
	
	e.Start(":8000")
}
