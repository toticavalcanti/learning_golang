package main

import (
	"errors"
	"fmt"
	"os"
)

//custom error --> struct
type DivZero struct{}

func (myerr *DivZero) Error() string {
	return "Cannot divide by 0! (struct)"
}

//custom error --> New
func divide(x int, y int) (int, error) {
	if y == 0 {
		return -1, errors.New("Cannot divide by 0! (New)")
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

	answer, err := divide(5, 0)
	if err != nil {
		// Handle the error!
		fmt.Println(err)
	} else {
		// No errors!
		fmt.Println(answer)
	}

	myerr := &DivZero{}
	if myerr != nil {
		// Handle the error with custom error!
		fmt.Println(myerr)
	} else {
		// No errors!
		fmt.Println(answer)
	}
}
