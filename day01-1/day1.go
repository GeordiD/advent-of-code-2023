package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var nonAlphanumericRegex = regexp.MustCompile(`[^0-9]+`)

func FindNumber(str string) int {
	var allNumbers = nonAlphanumericRegex.ReplaceAllString(str, "")
	var firstNum = string(allNumbers[0])
	var lastNum = string(allNumbers[len(allNumbers)-1])
	result, _ := strconv.Atoi(firstNum + lastNum)
	return result
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	var sum = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += FindNumber(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum is ", sum)
}
