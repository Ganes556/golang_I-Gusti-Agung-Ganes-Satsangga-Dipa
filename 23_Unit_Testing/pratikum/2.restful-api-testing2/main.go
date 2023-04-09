package main

import (
	conf "github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/configs"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/routes"
)

func main() {
  conf.InitDB()
  e := routes.New()
  e.Logger.Fatal(e.Start(":8000"))
}