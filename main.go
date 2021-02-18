package main

import (
	"log"

	"github.com/Reywaltz/web_test/cmd/journal-api/handlers"
	"github.com/Reywaltz/web_test/internal/repository/queries"
	"github.com/Reywaltz/web_test/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := postgres.NewCfg("cfg")

	db, err := postgres.NewDb(cfg)
	if err != nil {
		log.Fatal("Error")
	}

	stdres := queries.NewRepository(db)

	handler := handlers.NewUserGroupHandler(stdres)

	router := gin.Default()

	handler.Route(router)

	router.Run()
}
