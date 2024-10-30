package main

import (
	"fmt"
	"github/fbdaf/go-postgres/middleware"
	"github/fbdaf/go-postgres/router"
	"log"
	"net/http"
)

func main() {
	db := middleware.CreateConnection()
	defer db.Close()
	middleware.AutoMigrate(db)

	r := router.Router()
	fmt.Println("Starting server on port 8888...")
	log.Fatal(http.ListenAndServe(":8888", r))
}
