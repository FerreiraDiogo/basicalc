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

	fmt.Println("==========BASICALC  V1.0.1==========")
	fmt.Println("Instructions: Type the first operand, press enter. Type the operand and press enter again. Then type the third operand and press enter",
		"Valid operand values are '+', '-', '*', '/' and '%'")

	_, err := fmt.Scan(&firstInput, &operator, &secondInput)
	if err != nil {
		panic(err)
	}
	calculatedValue, calculationError := calculate(&firstInput, &operator, &secondInput)
	if calculationError != nil {
		panic(calculationError)
	}
	fmt.Println(fmt.Sprintf("The Calculated value is %f", calculatedValue))

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
		return 0, errors.New("No valid operand was found")
	}

}
