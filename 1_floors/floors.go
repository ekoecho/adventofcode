package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	floor := 0
	pos := 0
	log.Println("calculating floor")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewReader(file)

	for {
		if c, _, err := scanner.ReadRune(); err != nil {
			if err == io.EOF {
				log.Println(floor)
				break
			} else {
				log.Fatal(err)
			}

		} else {
			if floor == -1 {
				log.Println(pos)
				break
			}
			if c == 40 {
				floor = floor + 1
			}
			if c == 41 {
				floor = floor - 1
			}
			pos = pos + 1
		}

	}

}
