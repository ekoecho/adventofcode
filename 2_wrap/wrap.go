package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	wrap := 0
	ribbon := 0

	for scanner.Scan() {
		dim := strings.Split(scanner.Text(), "x")
		w, _ := strconv.Atoi(dim[0])
		l, _ := strconv.Atoi(dim[1])
		h, _ := strconv.Atoi(dim[2])

		dims := []int{}
		dims = append(dims, w)
		dims = append(dims, l)
		dims = append(dims, h)

		sort.Ints(dims)

		sides := []int{}
		sides = append(sides, w*l)
		sides = append(sides, w*h)
		sides = append(sides, l*h)

		sort.Ints(sides)

		wrap = wrap + sides[0] + (2 * sides[0]) + (2 * sides[1]) + (2 * sides[2])
		ribbon = ribbon + (dims[0] * 2) + (dims[1] * 2) + (dims[0] * dims[1] * dims[2])

	}
	log.Println(wrap)
	log.Println(ribbon)

}
