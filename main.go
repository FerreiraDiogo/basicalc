package main

import (
	"errors"
	"fmt"
)

func main() {

	var firstInput int
	var operator string
	var secondInput int

	fmt.Println("==========BASICALC==========")
	fmt.Println("Instructions: Type the first operand, press enter. Type the operand and press enter again. Then type the third operand and press enter",
		"Valid operand values are '+', '-', '*', '/' and '%'")

	_, err := fmt.Scan(&firstInput, &operator, &secondInput)
	if err != nil {
		fmt.Println("An input error has occurred: ", err)
	}
	calculatedValue, calculationError := calculate(&firstInput, &operator, &secondInput)
	if calculationError != nil {
		fmt.Println(calculationError)
	}
	fmt.Println(fmt.Sprintf("The Calculated value is %d", calculatedValue))

}

func calculate(firstInput *int, operator *string, secondInput *int) (int, error) {

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
		return *firstInput % *secondInput, nil
	default:
		return 0, errors.New("No valid operand was found")
	}

}
