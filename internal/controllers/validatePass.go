package controllers

import (
	"fmt"
	"log"
)

func ValidatePass() []byte {

	fmt.Println("Enter password")

	pss, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return []byte(pss)
}
