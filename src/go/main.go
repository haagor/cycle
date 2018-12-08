package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	l_router := mux.NewRouter()
	l_router.HandleFunc("/cycle/v1", day).Methods("POST")
	fmt.Println("listening...")
	log.Fatal(http.ListenAndServe(":8000", l_router))
}

type Day struct {
	Date    string
	Grade   int
	Comment string
}

func day(writer http.ResponseWriter, request *http.Request) {
	dayReceive, _ := ioutil.ReadAll(request.Body)
	var dayJson Day
	if err := json.Unmarshal(dayReceive, &dayJson); err != nil {
		fmt.Println("error:", err)
	}

	response := processDay(dayJson)

	json.NewEncoder(writer).Encode(response)
}

func processDay(day Day) string {
	fmt.Printf("day : %v, grade : %v, comment : \"%v\"\n", day.Date, day.Grade, day.Comment)
	return "Copy that ! See you tomorrow"
}
