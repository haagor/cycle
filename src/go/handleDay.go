package main

import (
	"encoding/json"
	"fmt"
)

const Index = "cyle"
const Mapping = `
{
	"mappings" : {
		"day" : {
			"properties" : {
				"date" : { "type" : "date" },
				"grade" : { "type" : "integer" },
				"good" : { "type" : "keyword" },
				"bad" : { "type" : "keyword" }
			}
		}
	}
}`

type Day struct {
	Date  string
	Grade int
	Good  string
	Bad   string
}

func dayToStringJson(day Day) string {
	d, err := json.Marshal(day)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(d)
}
