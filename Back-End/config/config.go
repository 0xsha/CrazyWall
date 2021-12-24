package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetKey ReadKey  func to get env value
func GetKey(key string) string {

	//_, ok := os.LookupEnv("debug")
	//if ok {

		// load .env file so it's accessible from GetEnv
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		return os.Getenv(key)
	//}

	return os.Getenv(key)
}
