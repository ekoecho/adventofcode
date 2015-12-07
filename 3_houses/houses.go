package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Location struct {
	X, Y int
}

func main() {

	routeMap := make(map[Location]int)
	roboRouteMap := make(map[Location]int)
	combinedRouteMap := make(map[Location]int)

	currentLoc := &Location{0, 0}
	roboCurrentLoc := &Location{0, 0}
	f, err := os.Open("input")

	if err != nil {
		log.Println(err)
	}

	input := bufio.NewReader(f)
	counter := 0

	routeMap[Location{currentLoc.X, currentLoc.Y}] = 1
	routeMap[Location{currentLoc.X, currentLoc.Y}] = 1
	for {
		if c, _, err := input.ReadRune(); err != nil {
			if err == io.EOF {
				log.Println("end of input")
				break
			}

		} else {
			if counter%2 == 0 {
				changeLoc(c, currentLoc)
				routeMap[Location{currentLoc.X, currentLoc.Y}] = 1
				combinedRouteMap[Location{currentLoc.X, currentLoc.Y}] = 1
			} else {
				changeLoc(c, roboCurrentLoc)
				roboRouteMap[Location{roboCurrentLoc.X, roboCurrentLoc.Y}] = 1
				combinedRouteMap[Location{roboCurrentLoc.X, roboCurrentLoc.Y}] = 1
			}
		}

		counter += 1

	}
	log.Println(len(combinedRouteMap))

}

func changeLoc(c rune, loc *Location) {
	switch string(c) {
	case "v":
		loc.Y = loc.Y - 1
	case "^":
		loc.Y = loc.Y + 1
	case "<":
		loc.X = loc.X - 1
	case ">":
		loc.X = loc.X + 1
	}

}
