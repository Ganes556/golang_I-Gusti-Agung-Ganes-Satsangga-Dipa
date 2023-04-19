package main

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/configs"
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/routes"
	_ "github.com/joho/godotenv/autoload"
)

func init(){
  configs.InitDB()
}
func main() {
  e := routes.New()
  e.Logger.Fatal(e.Start(":8000"))
}