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

	fmt.Println("==========BASICALC  V1.0.2==========")
	fmt.Println("Instructions: Type the first operand, press enter. Type the operand and press enter again. Then type the third operand and press enter",
		"Valid operator values are '+', '-', '*', '/' and '%'")

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
		if err == nil {
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
