package main

import (
	"errors"
	"fmt"
	"os"
)

//custom error --> struct
type DivZero struct{}

func (myerr *DivZero) Error() string {
	return "Cannot divide by 0! (Using Custom Errors)"
}

var myerr = &DivZero{}

// using builtin "errors" --> Using Builtin Errors
func divide_using_builtin(x int, y int) (int, error) {
	if y == 0 {
		return -1, errors.New("Cannot divide by 0! (Using Builtin Errors)")
	}
	return x / y, nil
}

// using custom DivZero "errors" --> Using Custom Errors
func divide_using_custom(x int, y int) (int, error) {
	if y == 0 {
		return -1, myerr
	}
	return x / y, nil
}

func main() {

	// basic use --> trying to open non existing file
	filename, err := os.Open("Nofile.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("filename", filename) //comes out as

	answer, err := divide_using_builtin(5, 0)
	if err != nil {
		// Handle the error!
		fmt.Println(err)
	} else {
		// No errors!
		fmt.Println(answer)
	}
	answer, err = divide_using_builtin(5, 0)
	if myerr != nil {
		// Handle the error with custom error!
		fmt.Println(myerr)
	} else {
		// No errors!
		fmt.Println(answer)
	}
}
