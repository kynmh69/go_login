package utils

import (
	"go_login/logging"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	logger := logging.GetLogger()
	envFilePath := ".env"
	err := godotenv.Load(envFilePath)
	if err != nil {
		logger.Fatalln("env load error.", err.Error())
	}
	logger.Debug("Load env.")
}
