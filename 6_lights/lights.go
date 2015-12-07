package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	grid := [1000][1000]bool{}
	count := 0

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		//log.Println(scanner.Text())
		grid = parseCommand(scanner.Text(), grid)
	}

	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if grid[x][y] == true {
				count++
			}
		}
	}
	log.Println(count)

}

func parseCommand(s string, g [1000][1000]bool) [1000][1000]bool {

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
				g[x][y] = true
			} else if action == "OFF" {
				g[x][y] = false
			} else {
				g[x][y] = !(g[x][y])
			}
		}
	}
	return g

}
