package main

import (
	"github.com/dsisconeto/arquitetura/api/handler"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler.ProductRegisterRouters(mux)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Fatal(http.ListenAndServe(":8080", mux))
}
