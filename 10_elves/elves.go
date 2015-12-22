package main

import (
	"log"
	"strconv"
)

func main() {

	//input := "1321131112"
	input := "112332"
	counter := 1
	prev := ""
	output := ""

	for i, c := range input {
		current := string(c)
		if i != 0 {
			if prev == current {
				counter++
			} else {
				returnstring := strconv.Itoa(counter) + prev
				output = output + returnstring
				counter = 1
			}
		}

		prev = current
	}
	log.Println(output)

}
