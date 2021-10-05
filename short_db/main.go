package short_db

import (
	"fmt"
	"math/rand"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type db interface {
	Setup()
	// ShortURLs Methods
	ReadShort(short string)
	WriteShort(url string)
	ListShort()
	// Stats Methods
	WriteStat(short string, IP string)
	ListStat()
}

type ShortURL struct {
	gorm.Model
	Shortcode string `json:"shortcode"`
	URL       string `json:"url"`
}

type Stat struct {
	gorm.Model
	Shortcode string
	IP        string
}

// ShortURLs Methods
func ReadShort(short string) string {
	db := GetDB()

	var shortResult ShortURL

	result := db.Find(&shortResult, "Shortcode = ?", short)

	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return shortResult.URL
}

func WriteShort(url string) ShortURL {
	db := GetDB()

	shortObject := ShortURL{
		Shortcode: GenerateShort(8),
		URL:       url,
	}

	if result := db.Create(&shortObject); result.Error != nil {
		fmt.Println(result.Error)
	}

	return shortObject
}

func ListShort() []ShortURL {
	db := GetDB()

	shortResults := []ShortURL{}

	if result := db.Find(&shortResults); result.Error != nil {
		fmt.Println(result.Error)
	}

	return shortResults
}

func WriteStat(short string, IP string) Stat {
	db := GetDB()

	statObject := Stat{
		Shortcode: short,
		IP:        IP,
	}

	if result := db.Create(&statObject); result.Error != nil {
		fmt.Println(result.Error)
	}

	return statObject
}

func ListStat() []Stat {
	db := GetDB()

	statResults := []Stat{}

	if result := db.Find(&statResults); result.Error != nil {
		fmt.Println(result.Error)
	}

	return statResults
}

// Private Helpers
func GetDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("short.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}

	return db
}

func Seed() {
	db := GetDB()

	db.Create(&ShortURL{
		Shortcode: "google",
		URL:       "https://google.com",
	})

	db.Create(&ShortURL{
		Shortcode: "yt",
		URL:       "https://youtube.com",
	})

	db.Create(&ShortURL{
		Shortcode: "spot",
		URL:       "https://spotify.com",
	})
}

func SetupDB() {
	db := GetDB()
	db.AutoMigrate(&ShortURL{}, &Stat{})

	Seed()
	// if _, err := os.Stat("short.db"); os.IsNotExist(err) {
	// }
}

// Helpers
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateShort(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
