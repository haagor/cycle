package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func injectFile() {
	deleteIndex(Index)
	creatIndex(Index, Mapping)

	file, _ := os.Open("/home/haag/go/src/github.com/cycle/src/go/days.txt")
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if line == nil {
			break
		}
		check(err)

		d, _ := parse(string(line))
		_ = processDay(d)
	}
}

func parse(line string) (Day, error) {
	fields := strings.Split(line, " : ")
	var d Day
	if len(fields) != 4 {
		return d, errors.New("strange line !")
	}
	d.Date = fields[0]
	d.Grade, _ = strconv.Atoi(fields[1])
	d.Good = strings.Split(fields[2], ", ")
	d.Bad = strings.Split(fields[3], ", ")

	return d, nil
}
