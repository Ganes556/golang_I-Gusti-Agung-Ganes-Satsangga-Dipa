package main

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/routes"
)

func main() {
  e := routes.New()
  e.Logger.Fatal(e.Start(":8000"))
}