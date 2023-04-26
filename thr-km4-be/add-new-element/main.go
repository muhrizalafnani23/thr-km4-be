package main

func addNewElement(array []interface{}, element interface{}, position string) []interface{} {
	if position == "up" {
		return append([]interface{}{element}, array...)
	} else if position == "down" {
		return append(array, element)
	} else {
		return array
	}
}

func main() {}
