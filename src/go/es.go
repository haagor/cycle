package main

import (
	"fmt"
	"net/http"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println("error:", err)
	}
}

func creatIndex(esIndex string, esMapping string) {
	request, err := http.NewRequest("PUT", "http://localhost:9200/"+esIndex, strings.NewReader(esMapping))
	check(err)
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	check(err)
	defer response.Body.Close()
}

func deleteIndex(esIndex string) {
	request, err := http.NewRequest("DELETE", "http://localhost:9200/"+esIndex, nil)
	check(err)
	response, err := http.DefaultClient.Do(request)
	check(err)
	defer response.Body.Close()
}

func sendDoc(esIndex string, esType string, doc string) {
	request, err := http.NewRequest("POST", "http://localhost:9200/"+esIndex+"/"+esType+"/", strings.NewReader(doc))
	check(err)
	request.Header.Set("Content-Type", "application/json")

	l_resp, err := http.DefaultClient.Do(request)
	check(err)
	defer l_resp.Body.Close()
}
