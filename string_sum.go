package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(inp string) (output string, err error) {
	input := strings.TrimSpace(inp)
	if input == "" {
		return "", fmt.Errorf("received empty, %w", errorEmptyInput)
	}
	var firstOperand, secondOperand, result int
	var operator rune
	var firstOperatorNegative, firstOperatorFull, isSecondOperatorFull bool
	if input[0] == '-' {
		firstOperatorNegative = true
	}
	input = strings.Trim(input, "-")
	for _, charCode := range input {
		switch {
		case charCode == ' ':
			continue
		case charCode == '+', charCode == '-':
			if firstOperatorFull {
				return "", fmt.Errorf("пустая строка, %w", errorNotTwoOperands)
			}
			if firstOperatorNegative {
				result = -firstOperand
			} else {
				result = firstOperand
			}
			operator = charCode
			firstOperatorFull = true
		default:
			count, err := strconv.Atoi(string(charCode))
			if err != nil {
				return "", fmt.Errorf("bad token, %w", err)
			}
			if !firstOperatorFull {
				firstOperand = firstOperand*10 + count
			} else {
				secondOperand = secondOperand*10 + count
				isSecondOperatorFull = true
			}
		}
	}
	if !isSecondOperatorFull {
		return "", fmt.Errorf("received less two operands, %w", errorNotTwoOperands)
	}
	if operator == '+' {
		return strconv.Itoa(result + secondOperand), nil
	} else {
		return strconv.Itoa(result - secondOperand), nil
	}
}
