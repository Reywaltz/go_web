package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Reywaltz/web_test/cmd/journal-api/handlers"
	"github.com/Reywaltz/web_test/internal/models/bus"
)

const (
	brigade     = "22"
	busName     = "dean"
	initURL     = "https://up-lab1.mirea.ru/bus"
	initCom     = "INIT_INSTANCE"
	updateCom   = "UPDATE_SUBSCRIPTION"
	addRowCom   = "ADD_ROW"
	contentType = "application/json;charset=UTF-8"
	getURL      = "https://up-lab1.mirea.ru/bus?to=%s&after=%s"
)

var (
	errBadStatusCode = errors.New("bad status code")
)

func initBus() ([]byte, error) {
	newBus := bus.Bus{
		From:    brigade,
		To:      busName,
		Subject: initCom,
		Data:    nil,
	}

	jsonData, err := json.Marshal(newBus)
	if err != nil {
		return nil, fmt.Errorf("%w: can't marshal json", err)
	}
	return jsonData, nil
}

func initPost(jsonData []byte) error {
	resp, err := http.Post(initURL, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("%w: post error", err)
	}

	if resp.StatusCode != http.StatusOK {
		return errBadStatusCode
	}
	fmt.Println("init request sent")
	return nil
}

func updateBus() ([]byte, error) {
	data := fmt.Sprintf(`{"address":"%s","entityName":"student","type":"COMMON"}`, brigade)

	newBus := bus.Bus{
		From:    brigade,
		To:      busName,
		Subject: updateCom,
		Data:    &data,
	}

	jsonData, err := json.Marshal(newBus)
	if err != nil {
		return nil, fmt.Errorf("%w: can't marshal json", err)
	}
	return jsonData, nil
}

func updateSubscription(jsonData []byte) error {
	resp, err := http.Post(initURL, contentType, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("%w: post error", err)
	}

	if resp.StatusCode != http.StatusOK {
		return errBadStatusCode
	}
	fmt.Println("update request sent")
	return nil
}

func pushStudents(studentHandler *handlers.StudentHandlers) error {

	students, err := studentHandler.StudentStorage.Students()
	if err != nil {
		return fmt.Errorf("%w can't get students from db", err)
	}

	for _, student := range students {
		data := fmt.Sprintf(`{"isBinariesChanged":false,"entityName":"student","plainData":{"id":%s,"surname":"%s","name":"%s","second_name":"%s","study_group_id":%s,"study_group":null},"binaryLinks":{}}`,
			strconv.Itoa(student.ID), student.Surname, student.Name, student.SecondName, strconv.Itoa(student.StudyGroupID))

		newBus := bus.Bus{
			From:    brigade,
			To:      busName,
			Subject: addRowCom,
			Data:    &data,
		}

		jsonData, err := json.Marshal(newBus)
		if err != nil {
			return fmt.Errorf("%w: can't marshal json", err)
		}
		fmt.Println(bytes.NewBuffer(jsonData))
		resp, err := http.Post(initURL, contentType, bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("%w: post error", err)
		}

		if resp.StatusCode != http.StatusOK {
			return errBadStatusCode
		}
	}
	fmt.Println("student pushed")
	return nil
}

func fetchBus() error {
	ticker := time.NewTicker(3 * time.Second)
	for range ticker.C {
		timestamp := time.Now().UnixNano() / 1000000
		resp, err := http.Get(fmt.Sprintf(getURL, brigade, strconv.FormatInt(timestamp, 10)))
		if err != nil {
			return fmt.Errorf("%w can't get info", err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("%w can't read body", err)
		}
		defer resp.Body.Close()
		fmt.Println(string(body))
	}
	return nil
}
