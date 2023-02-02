package main

import (
	"fmt"
	"gopkg.in/validator.v2"
)

type User struct {
	Edad string `validate:"min=18,max=100"`
}

func main() {
	user := User{"17"}

	if err := validator.Validate(user); err != nil {
		fmt.Println("Error", err)
	}
}
