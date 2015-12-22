package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("input")
	reader := bufio.NewReader(file)

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")

		speed, _ := strconv.Atoi(text[3])
		time, _ := strconv.Atoi(text[6])
		rest, _ := strconv.Atoi(text[13])
		resting := false
		runtime := 0
		restTime := 0
		distance := 0
		for i := 0; i <= 2503; i++ {
			//log.Println(speed, time, rest)
			if restTime == rest {
				resting = false
				restTime = 0
			}
			if resting {
				restTime++
			} else {
				runtime++
				distance = distance + speed
				if runtime == time {
					resting = true
				}

			}

		}
		log.Println(speed, time, rest)
		log.Println(distance)

	}

}
