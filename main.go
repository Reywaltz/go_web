package main

import (
	"log"

	"github.com/Reywaltz/web_test/cmd/journal-api/handlers"
	"github.com/Reywaltz/web_test/internal/repository/queries"
	"github.com/Reywaltz/web_test/pkg/postgres"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := postgres.NewCfg("cfg", "toml")

	db, err := postgres.NewDb(cfg)
	if err != nil {
		log.Fatal("Error")
	}

	repo := queries.New(db)

	groupHandler := handlers.NewStudyGroupHandler(repo)

	studentHanlder := handlers.NewStudentHandler(repo)

	journalHandler := handlers.NewJournalHandler(repo)

	router := gin.Default()

	groupHandler.Route(router)
	studentHanlder.Route(router)
	journalHandler.Route(router)

	router.Run()
}
