package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var numberRegex = regexp.MustCompile(`^([0-9]|one|two|three|four|five|six|seven|eight|nine|zero)`)

func GetNumberFromBeginningOfString(str string) string {
	var result = numberRegex.FindStringSubmatch(str)

	if len(result) == 0 {
		return ""
	} else {
		return result[0]
	}
}

func TransformWordToNumericChar(str string) string {
	switch str {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	case "zero":
		return "0"
	default:
		return str
	}
}

func FindNumber(str string) int {
	var firstNum = ""
	var lastNum = ""

	for i := range str {
		var result = GetNumberFromBeginningOfString(str[i:])
		if result != "" {
			if firstNum == "" {
				firstNum = TransformWordToNumericChar(result)
			}
			lastNum = TransformWordToNumericChar(result)
		}
	}

	result, _ := strconv.Atoi(firstNum + lastNum)
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
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
