package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	var firstInput float64
	var operator string
	var secondInput float64

	fmt.Println("==========BASICALC  V1.0.3==========")
	fmt.Println("Instructions: Just type the inputs as you are prompted", "Valid operator values are '+', '-', '*', '/' and '%', Type 'quit' to end execution")

	fmt.Println("Type the first operand: ")
	for {
		_, err := fmt.Scanf("%f", &firstInput)
		if err == nil {
			break
		} else {
			fmt.Println("You must type a numeric operand")
		}
	}

	fmt.Println("Type the operator: ")
	for {
		_, err := fmt.Scanf("%s", &operator)
		if err == nil && isValidOperand(operator) {
			break
		}
		fmt.Println("Invalid input.Valid operator values are '+', '-', '*', '/' and '%'")

	}

	fmt.Println("Type the second operand: ")
	for {
		_, err := fmt.Scanf("%f", &secondInput)
		if err == nil && isNotDivisionByZero(&operator, &secondInput) {
			break
		}
		fmt.Println("You must type a numeric operand")
	}

	calculatedValue, calculationError := calculate(&firstInput, &operator, &secondInput)
	if calculationError != nil {
		panic(calculationError)
	}
	fmt.Printf("The Calculated value is %f\n", calculatedValue)

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
