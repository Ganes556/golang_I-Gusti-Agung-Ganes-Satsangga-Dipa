package main

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/db"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/routes"
	"github.com/labstack/echo/v4"
)

func main() {
  e := echo.New()
	db.InitValidate(e)
	routes.Users(e)
  e.Logger.Fatal(e.Start(":8000"))
}