package app

import (
	"github.com/Ganes556/golang_I-Gusti-Agung-Ganes-Satsangga-Dipa/configs"
)

func Start() {
	configs.SetupEnv()
	db := configs.InitDB()
	router := startRoute(db)
	router.Start(":8000")	
}