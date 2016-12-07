package main

import (
	"D"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"product"
	"strings"

	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var db *gorm.DB
var category string

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
	db.AutoMigrate(&product.Lens{})
	db.AutoMigrate(&product.VideoCamera{})
}

// メーカーの取得
func getMaker(doc *goquery.Document) string {
	if maker := doc.Find(".makerLabel .cateBoxPare").Text(); maker != "" {
		return sjisToUtf8(doc.Find(".makerLabel .cateBoxPare").Text())
	}
	return sjisToUtf8(doc.Find(".makerLabel").Text())
}

// プロダクトの必要情報を取得
func getEachProductInfos(url string) {
	pinfo := new(product.PInfo)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal("failed to get product info doc")
	}
	pinfo.Maker = getMaker(doc)
	fmt.Println(pinfo.Maker)
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
		getEachProductInfos(detailURLs[i])
	}
}

// sjis to utf8
func sjisToUtf8(str string) string {
	str = strings.TrimSpace(str)
	output1, err := sjisToUtf8_1(str)
	if err != nil {
		output2, err := sjisToUtf8_2(output1)
		if err != nil {
			return str
		}
		return output2
	}
	return output1
}

// sjis to utf8 pattern 1
func sjisToUtf8_1(str string) (string, error) {
	output, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return str, err
	}
	return string(output), err
}

// sjis to utf8 pattern 2
func sjisToUtf8_2(str string) (string, error) {
	output, err := iconv.ConvertString(str, "shift-jis", "utf-8")
	if err != nil {
		return str, err
	}
	return string(output), err
}

// コマンドライン引数からURLを取得
func getFlagURL() string {
	flag.Parse()
	url := flag.Arg(0)
	urlParams := strings.Split(url, "/")
	if urlParams[2] != D.KakakuDomain || urlParams[5] != D.KakakuItemList {
		log.Fatal("url is not kakaku item list")
	}
	return url
}

// 商品のカテゴリーを取得
func setCategoryFromURL(url string) {
	urlParams := strings.Split(url, "/")
	category = urlParams[4]
}

func main() {
	EnvLoad()
	initDB()
	defer db.Close()
	url := getFlagURL()
	setCategoryFromURL(url)
	setProductInfos(url)
}
