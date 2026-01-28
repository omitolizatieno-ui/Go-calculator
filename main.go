package main

import (
	"os"
	"fmt"
	"strconv"
)

func main(){

	if len(os.Args) != 4 {
		fmt.Println("Error")
		return
	}

	num1, err1 := strconv.Atoi(os.Args[1])
	num2, err2 := strconv.Atoi(os.Args[3])
	operator := os.Args[2]

	if err1 != nil || err2 != nil {
		fmt.Println("Error")
		return
	}

	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		if num2 == 0 {
		fmt.Println("Error")
		return
		}
	case "%":
		if num2 == 0 {
		fmt.Println("Error")
		return
		}
	default:
		fmt.Println("Error")
	}
}
