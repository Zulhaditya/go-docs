package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if id != "ackxle" {
		return &notFoundError{Message: "data not found"}
	}

	// ok
	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		// terjadi error
		// penggunaan if else
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationErr.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundError.Error())
		// } else {
		// 	fmt.Println("unknown error", err.Error())
		// }

		// penggunaan switch case
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("unknown error")
		}
	} else {
		// sukses
		fmt.Println("sukses")
	}
}
