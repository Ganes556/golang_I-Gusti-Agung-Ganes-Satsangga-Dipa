package configs

import "github.com/joho/godotenv"

func SetupEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}