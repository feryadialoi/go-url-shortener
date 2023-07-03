package database

import (
	"fmt"
	"log"
	"testing"

	"github.com/feryadialoi/go-url-shortener/config"
)

func Test(t *testing.T) {
	db := MustNew(config.Config{
		DBHost: "localhost",
		DBPort: "5432",
		DBName: "url_shortener",
		DBUser: "postgres",
		DBPass: "password",
	})

	for i := 0; i < 1_000; i++ {
		sp := fmt.Sprintf("PATH%v", i+1)
		ru := fmt.Sprintf("http://localhost:8080/api/v1/url-shortens?shortPath=PATH%v", i+1)
		_, err := db.Exec("INSERT INTO url_shorten(short_path, real_url) VALUES ($1, $2)", sp, ru)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("done...")
}
