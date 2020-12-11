package io

import (
	"bufio"
	"log"
	"os"
)

// ReadFileLineByLine - reads file line by line in string slice
func ReadFileLineByLine(fileName string) (ids []string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 1
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

