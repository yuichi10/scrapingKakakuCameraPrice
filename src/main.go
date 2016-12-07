package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Err loading .env")
	}
}

func main() {
	EnvLoad()
	fmt.Print(os.Getenv("DB_USER"))
	fmt.Println("hello go")
}
