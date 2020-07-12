package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Get filename from command-line args
	fileName := os.Args[1:][0]
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	usePrefix := false
	counter := 1

	// Loop over each line in the file
	for scanner.Scan() {
		var prefix string
		if usePrefix {
			if scanner.Text() != "" {
				prefix = fmt.Sprintf("%s. ", strconv.Itoa(counter))
				counter++
			}
		}

		fmt.Println(prefix + scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
