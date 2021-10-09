package operations

import (
	"fmt"
	"strconv"
)

// Operations is for doing the calculations
func Operations(expr string) (int64, error) {
	var operatorSlice []string
	var operandSlice []int64
	for i := 0; i < len(expr); i++ {
		fmt.Printf("%c\n", expr[i])
		value, err := strconv.ParseInt(string(expr[i]), 0, 64)
		if err != nil {
			fmt.Println("The operator:", string(expr[i]))
			operatorSlice = append(operatorSlice, string(expr[i]))
			continue
		}
		fmt.Println("The integer is:", value)

		operandSlice = append(operandSlice, value)

	}
	fmt.Println("The integer array is:", operandSlice)
	fmt.Println("The operator array is:", operatorSlice)
	var finalValue int64
	for i := 0; i < len(operatorSlice); i++ {
		var operand1, operand2 int64
		if i == 0 {
			fmt.Println(operatorSlice[i], operandSlice[i], operandSlice[i+1])
			operand1 = operandSlice[i]
			operand2 = operandSlice[i+1]
		} else {
			fmt.Println(operatorSlice[i], operandSlice[i+1])
			operand1 = finalValue
			operand2 = operandSlice[i+1]
		}
		switch operatorSlice[i] {
		case "+":
			finalValue = operand1 + operand2
		case "-":
			finalValue = operand1 - operand2
		case "*":
			finalValue = operand1 * operand2
		case "/":
			if operand2 == 0 {
				return 0, fmt.Errorf("Error while dividing")
			}
			finalValue = operand1 / operand2
		}
	}
	return finalValue, nil
}
