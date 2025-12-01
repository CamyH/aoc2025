package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile() (*os.File, *bufio.Reader) {
	file, err := os.Open("day1/input.txt")

	reader := bufio.NewReader(file)

	if (err != nil) {
		fmt.Println(err)
		return nil, nil
	}

	return file, reader
}

func calculateNewRotation(rotation string, dial int) int {
	direction := string(rotation[0])
	number, err := strconv.Atoi(rotation[1:])

	if (err != nil) {
		fmt.Println(err)
		return -1
	}

	switch direction {
	case "L":
		dial = (dial - number) % 100
	case "R":
		dial = (dial + number) % 100
	default:
		fmt.Println("Invalid")
	}

	return dial
}

func Part1() int {
	var dialValue = 50
	var password = 0
	var file, reader = readFile()
	defer file.Close()
	
	for {
	line, _, err := reader.ReadLine()
		if len(line) > 0 {
			updatedDialValue := calculateNewRotation(string(line), dialValue)
			dialValue = updatedDialValue
			if (updatedDialValue == 0) {
				password += 1
			}
		}

		if err != nil {
			break
		}
	}

	return password
}