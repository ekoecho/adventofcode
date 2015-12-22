package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//grid := [1000][1000]bool{}
	newgrid := [1000][1000]int{}
	count := 0

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		//log.Println(scanner.Text())
		newgrid = parseCommand(scanner.Text(), newgrid)
	}

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			count = count + newgrid[x][y]
		}
	}
	log.Println(count)

}

func parseCommand(s string, g [1000][1000]int) [1000][1000]int {

	command := strings.Split(s, " ")
	var startPair string
	var endPair string
	var action string

	if command[0] == "toggle" {
		action = "TOGGLE"
		startPair = command[1]
		endPair = command[3]
	} else if command[1] == "on" {
		action = "ON"
		startPair = command[2]
		endPair = command[4]
	} else {
		action = "OFF"
		startPair = command[2]
		endPair = command[4]
	}

	startX, _ := strconv.Atoi(strings.Split(startPair, ",")[0])
	startY, _ := strconv.Atoi(strings.Split(startPair, ",")[1])
	endX, _ := strconv.Atoi(strings.Split(endPair, ",")[0])
	endY, _ := strconv.Atoi(strings.Split(endPair, ",")[1])

	log.Println(startX, startY, endX, endY, action)
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			if action == "ON" {
				g[x][y] = g[x][y] + 1
			} else if action == "OFF" {
				if g[x][y] > 0 {
					g[x][y] = g[x][y] - 1
				}
			} else {
				g[x][y] = g[x][y] + 2
			}
		}
	}
	return g

}
