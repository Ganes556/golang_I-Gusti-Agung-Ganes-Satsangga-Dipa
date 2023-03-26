package main

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/users"
	"github.com/labstack/echo/v4"
)

func main() {
  e := echo.New()
	users.Init(e)
  e.Logger.Fatal(e.Start(":8000"))
}