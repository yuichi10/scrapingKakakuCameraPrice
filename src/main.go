package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Err loading .env")
	}
}

func openDB() {
	var err error
	sql := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err = gorm.Open(os.Getenv("DB"), sql)
	if err != nil {
		fmt.Printf("エラー%v\n", err)
		return
	}
}

func main() {
	EnvLoad()
	openDB()
	fmt.Print(os.Getenv("DB_USER"))
	fmt.Println("hello go")
}
