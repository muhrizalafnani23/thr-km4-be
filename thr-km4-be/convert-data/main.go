package main

import (
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func convertData(data string) User {
	fields := strings.Split(data, ",")
	age, _ := strconv.Atoi(fields[1])

	return User{
		Name:    fields[0],
		Age:     age,
		Address: fields[2],
	}
}

func main() {}
