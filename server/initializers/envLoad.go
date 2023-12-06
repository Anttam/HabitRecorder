package initializers

import (
	"github.com/joho/godotenv"
	
	"log"
)

func LoadEnvVariables () {
	godotenv.Load()
	log.Println("Env Variables loaded")
}