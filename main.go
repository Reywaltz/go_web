package main

import (
	"log"
	"net/http"

	"github.com/Reywaltz/web_test/cmd/journal-api/handlers"
	"github.com/Reywaltz/web_test/internal/repository/queries"
	"github.com/Reywaltz/web_test/pkg/postgres"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := postgres.NewCfg("cfg", "toml")

	db, err := postgres.NewDb(cfg)
	if err != nil {
		log.Fatal("Can't establish conn to db", err)
	}

	repo := queries.New(db)

	groupHandler := handlers.NewStudyGroupHandler(repo)

	studentHanlder := handlers.NewStudentHandler(repo)

	journalHandler := handlers.NewJournalHandler(repo)

	subjectHandler := handlers.NewSubjectHandler(repo)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	router.StaticFS("journal/main/static", http.Dir("./web_test/static"))

	router.LoadHTMLFiles("web_test/index.html")

	groupHandler.Route(router)
	studentHanlder.Route(router)
	journalHandler.Route(router)
	subjectHandler.Route(router)

	router.Run()
}
