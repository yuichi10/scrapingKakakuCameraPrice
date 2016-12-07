package main

import (
	"fmt"
	"log"
	"os"
	"product"

	"github.com/PuerkitoBio/goquery"
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

func initDB() {
	openDB()
	db.DB()
	db.AutoMigrate(&product.DslrCamera{})
}

// 商品詳細へのリンクを取得
func getProductDetailURL(url string) []string {
	var urls []string
	urls = make([]string, 0)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Printf("getLinks Error: %v \n", err)
		return nil
	}
	doc.Find(".ckitemLink .ckitanker").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		urls = append(urls, string(url))
	})
	return urls
}

// プロダクトの情報を書き込む
func setProductInfos(url string) {
	detailURLs := getProductDetailURL(url)
	for i := 0; i < len(detailURLs); i++ {
		fmt.Println(detailURLs[i])
	}
}

func main() {
	EnvLoad()
	initDB()
	defer db.Close()
	setProductInfos("http://kakaku.com/camera/digital-slr-camera/itemlist.aspx")
}
