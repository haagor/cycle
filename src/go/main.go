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
	l_router.HandleFunc("/cycle/day", day).Methods("POST")
	l_router.HandleFunc("/cycle/init", initIndex).Methods("POST")
	l_router.HandleFunc("/cycle/erase", eraseIndex).Methods("POST")

	fmt.Println("listening...")
	log.Fatal(http.ListenAndServe(":8000", l_router))
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

func initIndex(writer http.ResponseWriter, request *http.Request) {
	creatIndex(Index, Mapping)
}

func eraseIndex(writer http.ResponseWriter, request *http.Request) {
	deleteIndex(Index)
}

func processDay(day Day) string {
	fmt.Printf("day : %v, grade : %v, good : %v, bad : %v\n", day.Date, day.Grade, day.Good, day.Bad)
	return "Copy that ! See you tomorrow"
}
