package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := 100
	output := reverse(input)
	fmt.Println(output)
}
func reverse(input int) int {
	var result int
	isNegative := false
	if input < 0 {
		isNegative = true
		input = input * -1
	}
	inputString := fmt.Sprint(input)
	outputString := ""
	inputStringArray := strings.Split(inputString, "")
	for i := len(inputStringArray) - 1; i >= 0; i-- {
		outputString = outputString + inputStringArray[i]
	}
	if isNegative == true {
		outputString = "-" + outputString
	}
	result, _ = strconv.Atoi(outputString)
	return result
}
