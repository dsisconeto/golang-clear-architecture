package main

import (
	"github.com/dsisconeto/arquitetura/api/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler.ProductRegisterRouters(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}



