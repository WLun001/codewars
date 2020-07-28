package main

import (
	"log"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	arg1 := os.Args[1]

	conn, err := net.Dial("tcp", "198.51.100.1:80")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	if _, err := conn.Write([]byte(arg1)); err != nil {
		log.Fatal(err)
	}

	rb := make([]byte, 128)
	bLen, err := conn.Read(rb)
	if err != nil {
		log.Fatal(err)
	}

	if bLen > 0 {
		byteV := rb[2:]
		splittedV := strings.Split(string(byteV), " ")
		var values []int
		for _, v := range splittedV {
			if i, err := strconv.Atoi(v); err != nil {
				log.Fatal(err)
			} else {
				values = append(values, i)
			}
		}

		sort.Ints(values)
		log.Println(values)
		return
	}

	// 259 23 7 14 13 10 12 8 55 1
}
