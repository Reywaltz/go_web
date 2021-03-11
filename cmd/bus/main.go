package main

import (
	"fmt"
	"log"

	"github.com/Reywaltz/web_test/cmd/journal-api/handlers"
	"github.com/Reywaltz/web_test/internal/repository/queries"
	"github.com/Reywaltz/web_test/pkg/postgres"
)

func main() {
	cfg, err := postgres.NewCfg("cfg", "toml")
	if err != nil {
		log.Fatal("Can't read or fing cfg file", err)
	}

	db, err := postgres.NewDB(cfg)
	if err != nil {
		log.Fatal("Can't establish conn to db", err)
	}

	repo := queries.New(db)

	studentHanlder := handlers.NewStudentHandler(repo)

	newBus, err := initBus()
	if err != nil {
		log.Fatal("can't create Bus")
	}
	if err = initPost(newBus); err != nil {
		fmt.Println("can't send post: ", err)
	}

	updateBus, err := updateBus()
	if err != nil {
		fmt.Println("%w can't create payload: ", err)
	}

	if err = updateSubscription(updateBus); err != nil {
		fmt.Println("can't send update request: ", err)
	}

	if err = pushStudents(studentHanlder); err != nil {
		fmt.Println("can't send post to push student", err)
	}

	if err := fetchBus(); err != nil {
		fmt.Println("can't get fetch data", err)
	}
}
