package main

func removeElement(array []interface{}, position string) []interface{} {
	if position == "left" {
		return array[1:]
	} else if position == "right" {
		return array[:len(array)-1]
	}
	return array
}

func main() {}
