package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(filepath string) {
	var envFilePath string

	if filepath == "" {
		envFilePath = ".env"
	} else {
		envFilePath = filepath
	}

	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalln("env load error.", err.Error())
	}
	log.Println("Load env.")
}
