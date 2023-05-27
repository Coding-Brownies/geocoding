package main

import (
	"fmt"
	"os"

	// import godotenv
	"github.com/joho/godotenv"
)

func main() {
	// load the .env file from the current directory
	godotenv.Load()

	// get the env variable
	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		fmt.Println("Error: env not found")
		return
	}

	// use the env variable
	fmt.Println("Content:", apiKey)
}
