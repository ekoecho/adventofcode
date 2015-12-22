package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	file, _ := os.Open("input")

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	total := 0
	memtotal := 0

	for scanner.Scan() {
		s := scanner.Text()
		log.Println(s)
		total += len(s)

		//log.Println(len(s))
		s = s[1 : len(s)-1]
		log.Println(s, (len(s)))
		rp := regexp.MustCompile("\\\\x..")
		s = rp.ReplaceAllString(s, "@")
		//log.Println(s)
		s = strings.Replace(s, "\\\\", "@", -1)
		s = strings.Replace(s, "\\\"", "@", -1)
		log.Println(s, len(s))
		memtotal += len(s)

	}

	log.Println("TOTAL:", total)
	log.Println("MEMTOTAL:", memtotal)
	log.Println(total - memtotal)

}
