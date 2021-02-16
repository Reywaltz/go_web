package main

import (
	"fmt"
	"log"

	"github.com/Reywaltz/web_test/cmd/journal-api/handlers"
	"github.com/Reywaltz/web_test/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := postgres.NewCfg("cfg")

	db, err := postgres.NewDb(cfg)
	if err != nil {
		log.Fatal("Error")
	}

	fmt.Println(db)

	router := gin.Default()
	router.GET("/", handlers.Mainhandler)
	router.Run()
}
