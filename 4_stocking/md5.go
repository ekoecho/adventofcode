package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"strconv"
)

func main() {

	i := 0

	for {

		hash := md5.New()

		//io.WriteString(hash, "abcdef609043")

		io.WriteString(hash, "bgvyzdsv"+strconv.Itoa(i))

		result := hex.EncodeToString(hash.Sum(nil))
		if result[:6] == "000000" {
			log.Println(result)
			log.Println(i)
			break
		}

		i++
	}

}
