package external

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func readEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}
	fmt.Printf("os.Getenv(\"MESSAGE\"): %v\n", os.Getenv("MESSAGE"))
}
