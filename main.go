package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	var firstInput string
	var operator string
	var secondInput string
	var controlFlag bool = true

	fmt.Println("==========BASICALC  V1.4.0==========")
	fmt.Println("Instructions: Just type the inputs as you are prompted", "Valid operator values are '+', '-', '*', '/' and '%', Type 'quit' to end execution")

	for controlFlag {
		for controlFlag {
			fmt.Println("Type the first operand: ")
			_, err := fmt.Scanf("%s", &firstInput)
			if err == nil {
				break
			} else {
				fmt.Println("You must type a numeric operand")
			}
		}
		if !mustBreakExecution(firstInput) {
			fmt.Println("Type the operator: ")
			for controlFlag {
				_, err := fmt.Scanf("%s", &operator)
				if mustBreakExecution(operator) {
					controlFlag = false
					break
				}
				if err == nil && isValidOperand(operator) {
					break
				}
				fmt.Println("Invalid input.Valid operator values are '+', '-', '*', '/' and '%'")

			}
		} else {
			controlFlag = false
		}

		if !mustBreakExecution(operator) && controlFlag {
			fmt.Println("Type the second operand: ")
			for controlFlag {
				_, err := fmt.Scanf("%s", &secondInput)
				convertedSecondInput, _ := strconv.ParseFloat(secondInput, 64)
				if err == nil && isNotDivisionByZero(&operator, &convertedSecondInput) {
					break
				}
				fmt.Println("You must type a numeric operand")
			}
		} else {
			controlFlag = false
		}

		if !mustBreakExecution(secondInput) && controlFlag {
			v1, _ := strconv.ParseFloat(firstInput, 64)
			v2, _ := strconv.ParseFloat(secondInput, 64)

			calculatedValue, calculationError := calculate(&v1, &operator, &v2)
			if calculationError != nil {
				panic(calculationError)
			}
			fmt.Printf("The Calculated value is %f\n", calculatedValue)
		} else {
			controlFlag = false
		}

	}

}

func isNotDivisionByZero(operator *string, operand *float64) bool {

	return !(*operator == "/" && *operand == 0.0)
}

func calculate(firstInput *float64, operator *string, secondInput *float64) (float64, error) {

	switch *operator {
	case "+":
		return *firstInput + *secondInput, nil
	case "-":
		return *firstInput - *secondInput, nil
	case "*":
		return *firstInput * *secondInput, nil
	case "/":
		return *firstInput / *secondInput, nil
	case "%":
		return math.Mod(*firstInput, *secondInput), nil
	default:
		return 0, errors.New("no valid operand was found")
	}

}

func isValidOperand(operand string) bool {
	switch operand {
	case "+", "-", "*", "/", "%":
		return true
	}
	return false
}

func mustBreakExecution(input string) bool {
	return strings.EqualFold("quit", input)
}
