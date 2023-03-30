package initializers

import (
	"api-test/utils"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	utils.LogError(err,"fail to load env")
	
}
