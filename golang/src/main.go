package main

import (
	"fmt"

	"github.com/johngalimi/random/golang/models"
)

func main() {
	fmt.Println("Hello world")

	u := models.User{
		ID:        2,
		FirstName: "Joe",
		LastName:  "Schmo",
	}

	fmt.Println(u)
}
