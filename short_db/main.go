package short_db

import (
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

type ShortURLs struct {
	gorm.Model
	Shortcode string `gorm:"uniqueIndex"`
	URL       string
}

type Stats struct {
	Shortcode string
	IP        string
}

// ShortURLs Methods
func ReadShort(short string) string {

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

	db.Create(&ShortURLs{
		Shortcode: "google",
		URL:       "https://google.com",
	})

	db.Create(&ShortURLs{
		Shortcode: "yt",
		URL:       "https://youtube.com",
	})

	db.Create(&ShortURLs{
		Shortcode: "spot",
		URL:       "https://spotify.com",
	})
}

func Setup() {
	db := GetDB()
	db.AutoMigrate(&ShortURLs{}, &Stats{})
	Seed()
}
