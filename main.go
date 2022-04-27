package main

import (
	"fmt"

	"github.com/rajatguptarg/go-webservice/models"
)

func main() {

	u := models.User{
		ID:        2,
		FirstName: "Rajat",
		LastName:  "Gupta",
	}

	fmt.Println(u)

}
