package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	l_router := mux.NewRouter()
	l_router.HandleFunc("/cycle/day", day).Methods("POST")
	l_router.HandleFunc("/cycle/init", initIndex).Methods("POST")
	l_router.HandleFunc("/cycle/erase", eraseIndex).Methods("POST")

	go listenStdin()

	fmt.Println("> listening...")
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
	fmt.Println(`> Index "` + Index + `" created`)
}

func eraseIndex(writer http.ResponseWriter, request *http.Request) {
	deleteIndex(Index)
	fmt.Println(`> Index "` + Index + `" erased`)
}

func processDay(day Day) string {
	doc := dayToStringJson(day)
	sendDoc(Index, Type, doc)
	fmt.Println(`> Inject one more day`)
	return "Copy that ! See you tomorrow"
}

func listenStdin() {
	reader := bufio.NewReader(os.Stdin)
	for {
		cmd, _ := reader.ReadString('\n')
		if cmd == "file\n" {
			fmt.Println(`> Copy that ! Inject`)
			injectFile()
		}
	}
}
